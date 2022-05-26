package graph_core_objects

import (
	"fmt"
	gonumgraph "gonum.org/v1/gonum/graph"
)

func CreateNode(
	NodeUuid string,
	NodeDisplayName string,
	graph *Graphs) *Nodes {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("unexpected panic: %v", r)
			panic(r)
		}
	}()

	node := graph.NewNode()

	prev := len(
		gonumgraph.NodesOf(
			graph.Nodes()))

	if graph.Node(node.ID()) != nil {

		curr := graph.Nodes().Len()

		if curr != prev {
			fmt.Printf("NewNode mutated graph: prev graph order != curr graph order, %d != %d", prev, curr)
		}
		fmt.Printf("NewNode returned existing: %#v", node)
	}

	newNode := &Nodes{
		NodeUuid:        NodeUuid,
		NodeDisplayName: NodeDisplayName,
		Node:            node,
	}

	return newNode

}
