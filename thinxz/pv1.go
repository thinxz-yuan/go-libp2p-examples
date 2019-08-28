package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
)

// 使用默认设置启动libp2p节点的文件, 打印节点的侦听地址, 然后关闭节点
func main() {
	// create a background context
	ctx := context.Background()

	// 启动LibP2P节点
	node, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	// 数据LibP2P节点监听地址
	fmt.Println("Listen addresses:", node.Addrs())

	// 关闭节点
	if err := node.Close(); err != nil {
		panic(err)
	}
}
