package server

import "go.uber.org/zap"

type Config struct {
	Dir           string
	Name          string
	PeerAddress   string
	ServerAddress string
	Peers         map[string]string
	Logger        *zap.SugaredLogger
}

func Bootstrap(config *Config) *RaftServer {

}
