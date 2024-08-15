package server

import (
	"sync"

	pb "github.com/xlcbingo1999/goraft/pkg/raftpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Stream interface {
	// Send sends a message to the peer.
	Send(*pb.RaftMessage) error

	// Recv receives a message from the peer.
	Recv() (*pb.RaftMessage, error)
}

type Remote struct {
	address string
	conn    *grpc.ClientConn
	client  pb.RaftClient
	// TODO(xlc_todo_delete): ???
	clientStream pb.Raft_ConsensusClient
	serverStream pb.Raft_ConsensusServer
}

type Peer struct {
	mu     sync.Mutex
	wg     sync.WaitGroup
	id     uint64
	stream Stream  // grpc 双向流
	remote *Remote // 远端信息

	recvc  chan *pb.RaftMessage // 流式读取数据发送管道
	metric chan pb.MessageType  // 监控管道
	close  bool                 // 是否要关闭peer
	logger *zap.SugaredLogger
}

func NewPeer(id uint64, address string, recvc chan *pb.RaftMessage, metric chan pb.MessageType, logger *zap.SugaredLogger) *Peer {
	return &Peer{
		id:     id,
		remote: &Remote{address: address},
		recvc:  recvc,
		metric: metric,
		logger: logger,
	}
}
