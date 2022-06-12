package graph_core_objects

import "gonum.org/v1/gonum/graph"

type NodeSets struct {
	NodeSetUuid        string
	NodeSetDisplayName string
	graph.Nodes
}
