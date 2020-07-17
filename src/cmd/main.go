package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"net"
	"os"
	"raft-tt/src/lib"
	"time"
)

func main() {
	
	// 节点配置
	config := raft.DefaultConfig()
	config.LocalID = "1"
	config.Logger = hclog.New(&hclog.LoggerOptions{
		Name:   "raftNode-1",
		Level:  hclog.LevelFromString("DEBUG"),
		Output: os.Stderr,
	})
	
	// 日志存储
	dir, _ := os.Getwd()
	logStore, err := raftboltdb.NewBoltStore(dir + "/data/log_store.bolt")
	if err != nil {
		panic(err)
	}
	
	// 数据存储
	stableStore, err := raftboltdb.NewBoltStore(dir + "/data/stable_store.bolt")
	if err != nil {
		panic(err)
	}
	
	// 存储快照
	snapshotStore := raft.NewDiscardSnapshotStore()
	
	// 通讯方式
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
	
	// 通信方式
	transport, err := raft.NewTCPTransport(addr.String(), addr, 5, time.Second*10, os.Stdout)
	
	// 持久化存储
	fsm := &lib.MyFSM{}
	
	node, err := raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		panic(err)
	}
	
	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID: config.LocalID,
				Address: transport.LocalAddr(),
			},
		},
	}
	
	node.BootstrapCluster(configuration)
}
