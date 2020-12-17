package bully

import (
	"distributed-algorithms/utils"
	"fmt"
	"sync"
)

// NodeIDType is unique id for each node
type NodeIDType = int

// PriorityType type of priority
type PriorityType = int

// Node is instance of physical computer
// with unique id, priority, is leader (currently running), failed or not
type Node struct {
	NodeID            NodeIDType
	CoordinatorNodeID NodeIDType
	Priority          PriorityType
	IsCoordinator     bool
	IsFailed          bool
}

// CreateNode - create new node
func CreateNode(id NodeIDType, priority int) Node {
	return Node{
		id,
		-1,
		priority,
		false,
		false,
	}
}

// Network is collection of node
// where each node is able to discover all other nodes
type Network struct {
	sync.WaitGroup
	sync.Mutex
	Nodes []Node
	Stop  chan bool
}

// CreateNetwork - init network
func CreateNetwork() Network {
	return Network{
		Nodes: make([]Node, 0),
		Stop:  make(chan bool),
	}
}

// MakeCoordinator - will make one of node coordinator
func (n *Network) MakeCoordinator(nodeID NodeIDType) {
	utils.Debug(fmt.Sprintf("Making Node %v coordinator", nodeID))
	for idx := range n.Nodes {
		if n.Nodes[idx].IsFailed {
			continue
		}
		n.Nodes[idx].CoordinatorNodeID = nodeID
		if n.Nodes[idx].NodeID == nodeID {
			n.Nodes[idx].IsCoordinator = true
		} else if n.Nodes[idx].IsCoordinator {
			n.Nodes[idx].IsCoordinator = false
		}
	}
}

// InsertNode - insert node in network
// insertion sort based on node priority
func (n *Network) InsertNode(node Node) {
	totalNodes := len(n.Nodes)
	if totalNodes == 0 {
		n.Nodes = []Node{node}
		n.MakeCoordinator(node.NodeID)
		return
	}

	n.Nodes = append(n.Nodes, node)
	itr := len(n.Nodes) - 2
	for itr >= 0 && n.Nodes[itr].Priority < node.Priority {
		n.Nodes[itr+1] = n.Nodes[itr]
		itr--
	}
	n.Nodes[itr+1] = node
	n.MakeCoordinator(n.Nodes[0].NodeID)
}

// IsCoordinatorFailed - check the health of coordinator
func (n *Network) IsCoordinatorFailed() bool {
	for _, node := range n.Nodes {
		if node.IsCoordinator && node.IsFailed {
			return true
		}
	}
	return false
}
