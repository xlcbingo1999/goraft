package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/xlcbingo1999/goraft/pkg/utils"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Servers []string `yaml:"servers"`
}

var (
	configFile string
	addNode    string
	removeNode string
	testPut    int
)

func init() {
	flag.StringVar(&configFile, "f", "config.yaml", "config file: -f config.yaml")
	flag.StringVar(&addNode, "add", "", "add node: -add raft_4=localhost:9204")
	flag.StringVar(&removeNode, "remove", "", "remove node:-remove raft_4=localhost:9204")
	flag.IntVar(&testPut, "test", 0, "test put: -test=10000")
}

func main() {
	flag.Parse()

	conf, err := os.ReadFile(configFile)
	if err != nil {
		log.Panicf("failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(conf, &config)
	if err != nil {
		log.Panicf("failed to unmarshal config file: %v", err)
	}

	// logger using for client
	logger := utils.GetLogger("client")
	// Sugar（语法糖的糖），只要一点点额外的性能损失（但是仍比大部分库快），可以比较简单地格式化输出。
	sugar := logger.Sugar()

	sugar.Infof("servers: %v", config.Servers)

	// 创建客户端
	c := client.NewClient(config.Servers, sugar)

	// 链接客户端
	c.Connect()

	// 增加节点
	if addNode != "" {
		nodes := make(map[string]string)
		for _, v := range strings.Split(addNode, ";") {
			node := strings.Split(v, "=")
			nodes[node[0]] = node[1]
		}
		// 增加到client中
		c.AddNode(nodes)
		return
	}

	// 删除节点
	if removeNode != "" {
		nodes := make(map[string]string)
		for _, v := range strings.Split(removeNode, ";") {
			node := strings.Split(v, "=")
			nodes[node[0]] = node[1]
		}
		// 删除到client中
		c.RemoveNode(nodes)
		return
	}

	// 测试
	if testPut > 0 {
		for i := 0; i < testPut; i++ {
			key := utils.RandStrihngBytesRmndr(rand.Intn(10) + 10)
			value := utils.RandStrihngBytesRmndr(20)
			c.Put(string(key), string(value))
		}
	}
}
