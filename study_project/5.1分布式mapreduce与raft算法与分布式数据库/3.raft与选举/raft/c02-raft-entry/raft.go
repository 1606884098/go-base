package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//

import (
	"math/rand"
	"sync"
	"time"
)
import "eth-1804/6.824/src/labrpc"

// import "bytes"
// import "encoding/gob"

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make().
//
type ApplyMsg struct {
	Index       int
	Command     interface{}
	UseSnapshot bool   // ignore for lab2; only used in lab3
	Snapshot    []byte // ignore for lab2; only used in lab3
}

//
// A Go object implementing a single Raft peer.
//
type Raft struct {
	mu        sync.Mutex          // Lock to protect shared access to this peer's state
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *Persister          // Object to hold this peer's persisted state
	me        int                 // this peer's index into peers[]
	// 自定义属性
	isLeader         bool          // 查看当前节点是否是一个leader
	resetTimer       chan struct{} // 重置选举超时
	electionTimer    *time.Timer   // 选举超时实例
	electionTimeout  time.Duration // 选举超时时间
	hearBeatInterval time.Duration // 心跳超时时间
	// Your data here (2A, 2B, 2C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.
	// Persistent state on all servers:需要持久的状态变量
	CurrentTerm  int          // 节点当前任期号
	VotedFor     int          // 投票给谁
	Logs         []LogEntry   // 日志结构
	commitCond   *sync.Cond   // 主要用于提交日志索引更新通知
	newEntryCond []*sync.Cond // 唤醒每个节点一致性检查
	commitIndex  int          // 已知的最大的已经被提交的日志条目的索引值
	lastApplied  int          // 最后被应用到状态机的日志条目索引值

	nextIndex []int // 对于每一个服务器，需要发送给他的下一个日志
	// 条目的索引值（初始化为领导人最后索引值加一）
	matchIndex []int // 对于每一个服务器，已经复制给他的日志的最高索引值

	applyCh  chan ApplyMsg // 日志提交成功之后，向服务发起的应用消息
	shutdown chan struct{} // 中断标志
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {

	var term int
	var isleader bool
	// Your code here (2A).
	rf.mu.Lock()
	defer rf.mu.Unlock()
	term = rf.CurrentTerm
	isleader = rf.isLeader
	return term, isleader
}

//
// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//
func (rf *Raft) Start(command interface{}) (int, int, bool) {
	index := -1
	term := 0
	isLeader := false // 假定当前不是leader
	select {
	case <-rf.shutdown:
		return index, term, isLeader
	default:

	}
	// Your code here (2B).
	rf.mu.Lock()
	defer rf.mu.Unlock()
	// 判断当前节点是否是leader，如果是leader,从client中添加日志
	if rf.isLeader {
		log := LogEntry{rf.CurrentTerm, command}
		rf.Logs = append(rf.Logs, log)
		index = len(rf.Logs) - 1 // 获取当前最后的日志索引
		term = rf.CurrentTerm
		isLeader = true

		rf.nextIndex[rf.me] = index + 1
		rf.matchIndex[rf.me] = index

		// client日志添加到leader成功
		rf.wakeupConsistencyCheck()
	}
	return index, term, isLeader
}

//
// the tester calls Kill() when a Raft instance won't
// be needed again. you are not required to do anything
// in Kill(), but it might be convenient to (for example)
// turn off debug output from this instance.
//
func (rf *Raft) Kill() {
	// Your code here, if desired.
}

//
// the service or tester wants to create a Raft server. the ports
// of all the Raft servers (including this one) are in peers[]. this
// server's port is peers[me]. all the servers' peers[] arrays
// have the same order. persister is a place for this server to
// save its persistent state, and also initially holds the most
// recent saved state, if any. applyCh is a channel on which the
// tester or service expects Raft to send ApplyMsg messages.
// Make() must return quickly, so it should start goroutines
// for any long-running work.
// 创建一个raft节点的服务
func Make(peers []*labrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me
	rf.applyCh = applyCh

	// Your initialization code here (2A, 2B, 2C).
	// 所以节点初始化状态都是follower
	rf.isLeader = false
	rf.VotedFor = -1

	rf.nextIndex = make([]int, len(peers))
	rf.matchIndex = make([]int, len(peers))

	rf.Logs = make([]LogEntry, 1)
	// 给一个空的
	rf.Logs[0] = LogEntry{
		Term:    0,
		Command: nil,
	}
	rf.shutdown = make(chan struct{})
	for i := 0; i < len(peers); i++ {
		// 初始化为leader最后索引值加一
		// 也就是leader的日志长度
		rf.nextIndex[i] = len(rf.Logs)
	}
	// 初始化重置超时选举
	rf.resetTimer = make(chan struct{})
	// commitCh
	rf.commitCond = sync.NewCond(&rf.mu)
	rf.newEntryCond = make([]*sync.Cond, len(peers))
	for i := 0; i < len(peers); i++ {
		rf.newEntryCond[i] = sync.NewCond(&rf.mu)
	}
	// 设置选举超时时间 400~800ms(注意rand的意义)
	rf.electionTimeout = time.Millisecond * (400 +
		time.Duration(rand.Int63()%400))
	rf.electionTimer = time.NewTimer(rf.electionTimeout)

	// 设置心跳超时时间
	rf.hearBeatInterval = time.Millisecond * 200
	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())
	// 初始化完成之后，启动选举进程
	go rf.electionDaemon()
	return rf
}
