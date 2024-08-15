package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	pb "github.com/xlcbingo1999/goraft/pkg/clientpb"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	clientId uint64
	reqSeq   uint64 // 请求序列号
	servers  []string
	client   pb.GoRaftClient // 这是一个rpc客户端, 由proto生成
	logger   *zap.SugaredLogger
}

func (c *Client) sendRequest(req *pb.Request) {
	// TODO?
}

func (c *Client) Put(key, value string) error {
	kv := &pb.KvPair{
		Key:   []byte(key),
		Value: []byte(value),
	}
	c.reqSeq++
	resp, err := c.client.Put(context.Background(), &pb.Request{
		ClientId: c.clientId,
		Seq:      c.reqSeq,
		// 把kv压缩到put请求中, 提交上去
		Cmd: &pb.Command{
			OperateType: pb.Operate_CONFIG,
			Put:         &pb.PutCommand{Data: []*pb.KvPair{kv}},
		},
	})
	if err != nil {
		return err
	}

	if resp.Success {
		return nil
	} else {
		return errors.New("put kv data failed.")
	}
}

func (c *Client) connect(address string) (pb.GoRaftClient, error) {
	// grpc的链接opts
	var opts []grpc.DialOption
	// 目前使用的是不安全的链接
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, fmt.Errorf("grpc connect to %s failed %v", address, err)
	}
	client := pb.NewGoRaftClient(conn)

	// grpc需要先注册
	resp, err := client.Register(context.Background(), &pb.Auth{Token: "token"})
	if err != nil || !resp.Success {
		conn.Close()
		return nil, fmt.Errorf("register to %s failed %v and success status: %v", address, err, resp.Success)
	}

	c.clientId = resp.ClientId
	return client, nil
}

func (c *Client) Connect() {
	var delay time.Duration
	for c.client == nil {
		address := c.servers[rand.Intn(len(c.servers))]
		client, err := c.connect(address)
		if err != nil {
			c.logger.Errorf("connect to %s failed: %v", address, err)
		} else if client != nil {
			c.client = client // 选择最新的client
			c.logger.Infof("connect to %s success.", address)
		}

		if c.client == nil {
			delay++
			if delay > 100 {
				delay = 0
			}
			// 每次失败都会睡眠，最多休眠10秒
			time.Sleep((delay/10 + 1) * time.Second)
		}
	}
}

func NewClient(servers []string, logger *zap.SugaredLogger) *Client {
	return &Client{
		servers: servers,
		logger:  logger,
	}
}
