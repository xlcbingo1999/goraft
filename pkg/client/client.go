package client

import "go.uber.org/zap"

type Client struct {
	clientId uint64
	reqSeq   uint64
	servers  []string
	client   pv.KvdbClient
	logger   *zap.SugaredLogger
}
