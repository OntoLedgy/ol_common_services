package graph_core_objects

import (
	"gonum.org/v1/gonum/graph"
)

type Edges struct {
	EdgeUuids string
	EdgeName  string
	EdgeEnd1  *Nodes
	EdgeEnd2  *Nodes
	graph.Edge
}
