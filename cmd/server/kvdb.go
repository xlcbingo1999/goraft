package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"path"

	"github.com/xlcbingo1999/goraft/pkg/utils"
	"gopkg.in/yaml.v2"
)

type Config struct {
	NodeConf *NodeConfig `yaml:"node"`
}

type NodeConfig struct {
	WorkDir       string            `yaml:"workDir"`
	Name          string            `yaml:"name"`
	PeerAddress   string            `yaml:"peerAddress"`
	ServerAddress string            `yaml:"serverAddress"`
	Servers       map[string]string `yaml:"servers"`
}

var configFile string

func init() {
	flag.StringVar(&configFile, "f", "config.yaml", "config file")
}

func main() {
	flag.Parse()
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("read config file %s failed %v", conf, err)
	}

	var config Config
	err = yaml.Unmarshal(conf, &config)
	if err != nil {
		log.Panicf("unmarshal config file %s failed %v", configFile, err)
	}

	logger := utils.GetLogger(path.Join(config.NodeConf.WorkDir, config.NodeConf.Name))
	sugar := logger.Sugar()

	allServers, _ := json.Marshal(config.NodeConf.Servers)
	sugar.Infof("configFile: %s allServers: %v", configFile, allServers)

	server.Bootstrap(&server.Config{
		Dir:           config.NodeConf.WorkDir,
		Name:          config.NodeConf.Name,
		PeerAddress:   config.NodeConf.PeerAddress,
		ServerAddress: config.NodeConf.ServerAddress,
		Peers:         config.NodeConf.Servers,
		Logger:        sugar,
	}).Start()
}
