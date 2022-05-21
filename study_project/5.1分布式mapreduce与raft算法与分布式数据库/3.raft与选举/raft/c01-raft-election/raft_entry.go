package raft

// raft 日志复制管理
// 日志结构
type LogEntry struct {
	Term    int         // 任期号
	Command interface{} // client端发送的命令
}
