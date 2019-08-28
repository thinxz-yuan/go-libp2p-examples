package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
)

// 使用默认设置启动libp2p节点的文件, 打印节点的侦听地址, 然后关闭节点
func main() {
	// create a background context
	ctx := context.Background()

	// 启动LibP2P节点
	node, err := libp2p.New(
		ctx,
		// 配置节点的监听地址, listens 127.0.0.1 on TCP port 2000 on the IPv4
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"),
	)
	if err != nil {
		panic(err)
	}

	// 数据LibP2P节点监听地址
	fmt.Println("Listen addresses:", node.Addrs())

	// 等待关闭信号, SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")

	// 关闭节点
	if err := node.Close(); err != nil {
		panic(err)
	}
}
