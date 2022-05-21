package raft

// raft 日志复制管理
// 日志结构
type LogEntry struct {
	Term    int         // 任期号
	Command interface{} // client端发送的命令
}

// 日志请求结构
type AppendEntriesArgs struct {
	Term int // leader任期号
	// 在raft中，有可能会出现client直接连上follower,此时,follower需要给client
	// 发送leaderId，方便follower重定向
	LeaderId     int        // leader ID so follower can redirect clients
	PrevLogIndex int        // 新的日志条目紧随之前的索引值
	PrevLogTerm  int        // PrevLogIndex任期
	Entries      []LogEntry // 准备存储的日志条目
	leaderCommit int        // leader已经提交的日志索引
}

// 日志复制响应结构
type AppendEntriesReply struct {
	CurrentTerm int // 当前任期号，主要用于leader更新自己
	// 跟随者包含了匹配上 prevLogIndex
	// 和 prevLogTerm 的日志时为真
	Success bool

	// 自定义冲突相关变量
	ConfliceTerm int // 冲突日志的任期编号
	FirstIndex   int // 存储第一个冲突编号的日志索引
}

// 唤醒一致性检查
func (rf *Raft) wakeupConsistencyCheck() {
	for i := 0; i < len(rf.peers); i++ {
		if i != rf.me {
			rf.newEntryCond[i].Broadcast()
		}
	}
}

// 启动日志复制进程
func (rf *Raft) logEntryAgreeDaemon() {
	// 遍历节点，向其它每个节点发起日志复制操作
	for i := 0; i < len(rf.peers); i++ {
		if i != rf.me {
			go rf.consistencyCheckDaemon(i)
		}
	}
}

// 发起日志复制操作
func (rf *Raft) consistencyCheckDaemon(n int) {
	for {
		rf.mu.Lock()
		// 每个节点都在等待client提交命令到leader上去
		rf.newEntryCond[n].Wait()
		select {
		case <-rf.shutdown:
			rf.mu.Unlock()
			return
		default:
		}

		// 判断节点角色，只有leader才能发起日志复制
		if rf.isLeader {
			var args AppendEntriesArgs
			args.Term = rf.CurrentTerm
			args.LeaderId = rf.me
			args.leaderCommit = rf.commitIndex
			args.PrevLogIndex = rf.nextIndex[n] - 1
			args.PrevLogTerm = rf.Logs[args.PrevLogIndex].Term
			// 判断是否有新的日志进来
			// len(rf.Logs) : leader当前的日志总数
			// rf.nextIndex[n]:leader要发送给节点n的下一个日志索引
			// leader的日志长度大于leader所知道的follow n的日志长度
			if rf.nextIndex[n] < len(rf.Logs) {
				// 添加新的日志
				args.Entries = append(args.Entries, rf.Logs[rf.nextIndex[n]:]...)
			} else {
				args.Entries = nil
			}
			rf.mu.Unlock()

			replyCh := make(chan AppendEntriesReply, 1)
			go func() {
				var reply AppendEntriesReply
				// 发起日志复制请求
				if rf.sendAppendEntries(n, &args, &reply) {
					replyCh <- reply
				}
			}()

			// 获取响应

		}
	}
}

// 发起日志复制的请求
func (rf *Raft) sendAppendEntries(server int, args *AppendEntriesArgs, reply *AppendEntriesReply) bool {
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	return ok
}
