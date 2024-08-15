package server

import (
	"github.com/xlcbingo1999/goraft/pkg/clientpb"
	pb "github.com/xlcbingo1999/goraft/pkg/raftpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RaftServer struct {
	pb.RaftServer
	clientpb.GoRaftServer

	dir           string
	id            uint64
	name          string
	peerAddress   string
	serverAddress string

	raftServer *grpc.Server
	kvServer   *grpc.Server
	peers      map[uint64]*Peer
	tmpPeers   map[uint64]*Peer

	incomingChan chan *pb.RaftMessage

	encoding raft.Encoding
	node     *raft.RaftNode
	storage  *raft.RaftStorage

	cache       map[string]interface{}
	leaderLease int64 // leader有效期
	close       bool
	stopc       chan struct{}
	metric      chan pb.MessageType
	logger      *zap.SugaredLogger
}
