package raft

import (
	pb "github.com/xlcbingo1999/goraft/pkg/raftpb"
	"go.uber.org/zap"
)

const MAX_LOG_ENTRY_SEND = 1000

type RaftState int // raft节点类型

const (
	CANDIDATE_STATE RaftState = iota
	FOLLOWER_STATE
	LEADER_STATE
)

type Raft struct {
	id                    uint64
	state                 RaftState
	leader                uint64   // 只有一个, 保留当前节点看到的leader
	currentTerm           uint64   // 当前任期
	voteFor               uint64   // 投票给谁
	raftlog               *RaftLog // 日志
	cluster               *Cluster // 集群节点
	electionTimeout       int      // 选举超时时间
	heartbeatTimeout      int      // 心跳超时时间
	randomElectionTimeout int      // 随机选举超时时间

	electionTick  int                   // 选举计时器
	heartbeatTick int                   // 心跳计时器
	Tick          func()                // 计时器函数, leader是心跳时钟，其他是选取时钟
	handleMessage func(*pb.RaftMessage) // 消息处理函数
	Msg           []*pb.RaftMessage     // 等待发送的消息
	ReadIndex     []*ReadIndexPesp      // 检查leader完成的readindex
	logger        *zap.SugaredLogger
}

func NewRaft(id uint64, storage Storage, peers map[uint64]string, logger *zap.SugaredLogger) *Raft {

}
