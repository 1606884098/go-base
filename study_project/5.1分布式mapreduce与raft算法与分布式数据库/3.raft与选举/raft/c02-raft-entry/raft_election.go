package raft

import (
	"sync"
	"time"
)

// raft leader选举管理

//
// example RequestVote RPC arguments structure.
// field names must start with capital letters!
//
// 投票请求参数
type RequestVoteArgs struct {
	// Your data here (2A, 2B).
	Term         int // 候选人任期号
	CandidateId  int // 候选人的ID
	LastLogIndex int // 候选人最后的日志索引
	LastLogTerm  int // 候选人最后的日志任期号
}

//
// example RequestVote RPC reply structure.
// field names must start with capital letters!
//
type RequestVoteReply struct {
	// Your data here (2A).
	CurrentTerm  int  // 当前任期号
	VotedGranted bool // 是否获取到了该节点的投票
}

//
// example RequestVote RPC handler.
//
// 接收到投票请求之后，进行处理
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	// Your code here (2A, 2B).
	rf.mu.Lock()
	defer rf.mu.Unlock()
	// 当前节点的最后一个日志索引获取
	lastLogIdx := len(rf.Logs) - 1
	// 当前节点最后一个日志包含的任期号
	lastLogTerm := rf.Logs[lastLogIdx].Term
	// 判断leader的任期号与自己的任期号谁大
	if args.Term < rf.CurrentTerm {
		reply.CurrentTerm = rf.CurrentTerm
		reply.VotedGranted = false
	} else {
		if args.Term > rf.CurrentTerm {
			// 将当前节点状态变更为follower
			rf.CurrentTerm = args.Term
			rf.isLeader = false
			rf.VotedFor = -1
		}
		//  TODO 如果 当前节点的votedFor 为空或者为 candidateId，
		//  TODO 并且候选人的日志至少和自己一样新，
		// TODO 那么就投票给他
		if rf.VotedFor == -1 || rf.VotedFor == args.CandidateId {
			// 候选人的日志是否和自己的最后一个日志一样新
			if args.LastLogTerm == lastLogTerm &&
				args.LastLogIndex >= lastLogIdx ||
				args.LastLogTerm > lastLogTerm {
				rf.resetTimer <- struct{}{}
				// 更改自己这个节点的状态
				rf.isLeader = false
				rf.VotedFor = args.CandidateId
				// 投票
				reply.VotedGranted = true
			}
		}
	}
}

//
// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
//
func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) bool {
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
	return ok
}

// 启动选举进程
func (rf *Raft) electionDaemon() {
	for {
		select {
		// 接收到重置请求之后的处理
		case <-rf.resetTimer:
			if !rf.electionTimer.Stop() {
				// 发送超时
				<-rf.electionTimer.C
			}
			// 重置选举超时
			rf.electionTimer.Reset(rf.electionTimeout)
		case <-rf.electionTimer.C:
			// 超时，也就是说follower在指定时间内没有接收到来自
			// leader的信息，就自己变成 candidate
			// 向其它节点发起投票请求
			go rf.canvassVotes()
			// 重置选举超时
			rf.electionTimer.Reset(rf.electionTimeout)
		}
	}
}

// 填充请求参数
func (rf *Raft) fillRequestVoteArgs(args *RequestVoteArgs) {
	// 修改状态，加锁保证数据安全
	rf.mu.Lock()
	defer rf.mu.Unlock()

	rf.CurrentTerm += 1 // 任期号加1
	rf.VotedFor = rf.me // 投给自己

	args.Term = rf.CurrentTerm
	args.CandidateId = rf.me
	args.LastLogIndex = len(rf.Logs) - 1
	args.LastLogTerm = rf.Logs[args.LastLogIndex].Term
}

// 发起选举请求
func (rf *Raft) canvassVotes() {
	// 请求参数
	var voteArgs RequestVoteArgs
	rf.fillRequestVoteArgs(&voteArgs)
	// 获取节点数量
	peers := len(rf.peers)
	// 设置缓存channel,大小为peers，保存结果
	replyCh := make(chan RequestVoteReply, peers)

	var wg sync.WaitGroup
	// 正式发起投票请求
	for i := 0; i < peers; i++ {
		if i == rf.me {
			rf.resetTimer <- struct{}{}
		} else {
			wg.Add(1)
			// 对每个节点发起投票
			go func(n int) {
				defer wg.Done()
				var reply RequestVoteReply

				// 投票RPC请求结果
				doneCh := make(chan bool, 1)
				go func() {
					ok := rf.sendRequestVote(n, &voteArgs, &reply)
					// 将请求结果传入doneCh中
					doneCh <- ok
				}()

				select {
				case ok := <-doneCh:
					if !ok {
						return
					}
					// 响应的投票结果传入replyCh
					replyCh <- reply
				}
			}(i)
		}
	}
	// 另起一个协程关闭结果通道
	go func() { wg.Wait(); close(replyCh) }()
	// 统计票数结果，自己给自己投一票，所以初始值为1
	var votes = 1
	// 遍历缓存通道，获取每一个响应中的投票结果
	for reply := range replyCh {
		if reply.VotedGranted == true {
			// 得到了当前返回的节点的票
			if votes++; votes > peers/2 {
				rf.mu.Lock()
				rf.isLeader = true // 成功当选Leader
				rf.mu.Unlock()
				// 重置相关状态
				rf.resetOnElection()
				// 发起心跳机制，防止其它的追随者变成候选人
				go rf.heartbeatDaemon()
				// 当选leader之后，发起日志复制操作
				go rf.logEntryAgreeDaemon()
				return
			}
		} else if reply.CurrentTerm > voteArgs.Term {
			// reply.CurrentTerm:follower的任期号
			// voteArgs.Term:准备竞选leader的candidate的任期号
			// 改变状态，重新回到follower的状态
			rf.mu.Lock()
			rf.isLeader = false
			rf.VotedFor = -1
			rf.CurrentTerm = reply.CurrentTerm
			rf.mu.Unlock()
			rf.resetTimer <- struct{}{}
			return
		}
	}
}

// 启动心跳进程
func (rf *Raft) heartbeatDaemon() {
	for {
		if _, isLeader := rf.GetState(); isLeader {
			// 只要是leader，就可以不断重置选举超时
			rf.resetTimer <- struct{}{}
		} else {
			break
		}
		// 设置心跳发送间隔
		time.Sleep(rf.hearBeatInterval)
	}
}

// 当一个领导人刚获得权力的时候，
// 他初始化所有其它节点的 nextIndex 值为自己的最后一条日志的index加1
func (rf *Raft) resetOnElection() {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	// 节点数量
	count := len(rf.peers)
	// 日志长度，恰好是最后一个日志的index+1
	length := len(rf.Logs)
	for i := 0; i < count; i++ {
		rf.nextIndex[i] = length
		// 对于每一个服务器，已经复制给他的日志的最高索引值
		rf.matchIndex[i] = 0
		if i == rf.me {
			// leader日志复制给自己
			rf.matchIndex[i] = length - 1
		}
	}
}
