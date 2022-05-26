package graph_core_objects

import (
	"gonum.org/v1/gonum/graph"
)

type Nodes struct {
	NodeUuid        string
	NodeDisplayName string
	graph.Node
}
