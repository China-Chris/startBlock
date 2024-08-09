package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"startblock/api"
	"startblock/blockchain"
	"startblock/p2p"
)

func main() {
	// 命令行参数解析
	isSeedNode := flag.Bool("seed", false, "是否启动为种子节点")
	address := flag.String("address", "/ip4/0.0.0.0/tcp/5000", "节点监听地址")
	httpAddress := flag.String("http-address", ":8081", "HTTP 服务器监听地址")
	flag.Parse()

	// 创建一个新的区块链
	bc := blockchain.NewBlockchain(10000)
	fmt.Println("区块链创建成功")

	// 种子节点列表
	seedNodes := []string{"/ip4/127.0.0.1/tcp/5000/p2p/QmSeedNodeID"} // 替换为种子节点的 Multiaddress 和 Peer ID

	// 创建节点
	node, err := p2p.NewNode(*address, seedNodes, *isSeedNode)
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}
	defer func() {
		if err := node.Host.Close(); err != nil {
			log.Printf("Failed to close node: %v", err)
		}
	}()

	// 创建 APIHandler
	apiHandler := api.NewAPIHandler(bc)
	fmt.Println("APIHandler 创建成功")

	// 设置路由
	router := api.SetupRouter(apiHandler)
	fmt.Println("路由设置成功")

	// 启动 HTTP 服务器
	fmt.Printf("Starting HTTP server on %s\n", *httpAddress)
	if err := http.ListenAndServe(*httpAddress, router); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
