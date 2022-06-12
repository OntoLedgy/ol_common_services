package graph_core_objects

import (
	"gonum.org/v1/gonum/graph/simple"
)

func CreateNewGraph(graphType string) *Graphs {

	switch graphType {
	case "directed":
		newSimpleGraph := simple.NewDirectedGraph()

		newGraph := new(Graphs)
		newGraph.GraphType = graphType
		newGraph.DirectedGraph = newSimpleGraph

		return newGraph

	case "undirected":
		newSimpleGraph := simple.NewUndirectedGraph()

		newGraph := new(Graphs)
		newGraph.GraphType = graphType
		newGraph.UndirectedGraph = newSimpleGraph

		return newGraph

	default:
		return nil

	}

}
