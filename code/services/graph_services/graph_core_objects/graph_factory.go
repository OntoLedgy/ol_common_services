package graph_core_objects

import (
	"gonum.org/v1/gonum/graph/simple"
)

func CreateNewGraph(graphType int) *Graphs {

	switch graphType {
	case 0:
		newSimpleGraph := simple.NewDirectedGraph()

		newGraph := new(Graphs)

		newGraph.DirectedGraph = newSimpleGraph

		return newGraph

	default:
		return nil

	}

}
