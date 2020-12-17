package bully

import (
	"fmt"

	"github.com/Marvin9/distributed-algorithms/utils"
)

// Start network simulation
func (n *Network) Start() {
	n.Add(1)
	go n.bully()
}

func (n *Network) ping(nodeID int) bool {
	for _, node := range n.Nodes {
		if node.NodeID == nodeID {
			if node.IsFailed {
				return false
			}
			return true
		}
	}
	return false
}

func (n *Network) election(nodeIndex int) {
	utils.Debug(fmt.Sprintf("Node %v is holding election", n.Nodes[nodeIndex].NodeID))
	itr := nodeIndex - 1
	for itr >= 0 {
		// to get the feeling of distribution, I intentionally implemented verbose ping
		OK := n.ping(n.Nodes[itr].NodeID)
		if OK {
			utils.Debug(fmt.Sprintf("Node %v with high priority is up.", n.Nodes[itr].NodeID))
			// it's now upto Node[itr]
			n.election(itr)
			return
		}
		itr--
	}

	// if no greater priority node are active
	n.MakeCoordinator(n.Nodes[nodeIndex].NodeID)
}

func (n *Network) bully() {
	defer n.Done()
	totalNodes := len(n.Nodes)
	i := 0

	for {
		n.Lock()
		utils.Debug(fmt.Sprintf("Node %v is in process", n.Nodes[i].NodeID))
		utils.Debug(n.Nodes)

		if n.IsCoordinatorFailed() {
			utils.Debug("Coordinator node failed")
			n.election(i)
		}

		i = (i + 1) % totalNodes
		n.Unlock()

		var in string
		utils.Debug("Press i for input mode, c for continue...")
		fmt.Scanf("%s", &in)
		switch in {
		case "i":
			utils.Debug(n.Nodes)
			n.Controll()
			utils.Debug(n.Nodes)
		case "s":
			continue
		}
	}
}

// Controll is used to make up and down nodes to feel the simulation and bully algorithm
func (n *Network) Controll() {
	var nodeID NodeIDType
	var operation int
	fmt.Printf("\nEnter node Id and 0/1 to take that node down/up: ")
	fmt.Scanf("%d %d", &nodeID, &operation)

	for idx := range n.Nodes {
		if n.Nodes[idx].NodeID == nodeID {
			if operation == 0 {
				n.Nodes[idx].IsFailed = true
			} else {
				n.Nodes[idx].IsFailed = false
				n.election(idx)
			}
			return
		}
	}
}
