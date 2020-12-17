package main

import (
	"github.com/Marvin9/distributed-algorithms/bully-algorithm"
)

func simulateBully() {
	network := bully.CreateNetwork()

	nodes := []bully.Node{
		bully.CreateNode(1, 1),
		bully.CreateNode(2, 2),
		bully.CreateNode(3, 3),
		bully.CreateNode(4, 4),
	}

	for _, node := range nodes {
		network.InsertNode(node)
	}

	network.Start()
	network.Wait()
}

func main() {
	simulateBully()
}
