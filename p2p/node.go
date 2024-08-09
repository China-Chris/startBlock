package p2p

import (
	"context"
	"fmt"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type Node struct {
	Host       *libp2p.Host
	PubSub     *pubsub.PubSub
	Topic      *pubsub.Topic
	Sub        *pubsub.Subscription
	IsSeedNode bool
	Address    string
}

func NewNode(address string, seedNodes []string, isSeedNode bool) (*Node, error) {
	ctx := context.Background()

	// 创建 libp2p 节点
	host, err := libp2p.New(libp2p.ListenAddrStrings(address))
	if err != nil {
		return nil, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	// 创建 PubSub
	ps, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		return nil, fmt.Errorf("failed to create GossipSub: %w", err)
	}

	// 加入主题
	topic, err := ps.Join("blockchain-topic")
	if err != nil {
		return nil, fmt.Errorf("failed to join topic: %w", err)
	}

	// 订阅主题
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	node := &Node{
		Host:       &host,
		PubSub:     ps,
		Topic:      topic,
		Sub:        sub,
		IsSeedNode: isSeedNode,
		Address:    address,
	}

	// 启动节点的消息处理
	go node.handleMessages(ctx)

	// 如果是普通节点，连接种子节点
	if !isSeedNode {
		for _, seed := range seedNodes {
			addrInfo, err := peer.AddrInfoFromString(seed)
			if err != nil {
				log.Printf("Failed to parse seed node address: %v", err)
				continue
			}
			if err := host.Connect(ctx, *addrInfo); err != nil {
				log.Printf("Failed to connect to seed node: %v", err)
			}
		}
	}

	return node, nil
}

func (n *Node) handleMessages(ctx context.Context) {
	for {
		msg, err := n.Sub.Next(ctx)
		if err != nil {
			log.Printf("Error reading from topic subscription: %v", err)
			continue
		}
		fmt.Printf("Received message: %s\n", msg.Data)
	}
}

func (n *Node) Start() error {
	// 节点启动逻辑，当前没有具体实现，可以根据需求添加
	return nil
}
