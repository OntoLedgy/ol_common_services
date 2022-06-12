package same_as_processing

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/graph_core_objects"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"
)

func ProcessSameAsLinks(edgesList graph.Edges) [][]graph.Node {

	simpleGraph := graph_core_objects.CreateNewGraph("undirected")
	edgesList.Reset()
	edgesList.Next()

	for edge := edgesList.Edge(); edge != nil; edgesList.Next() {

		simpleGraph.UndirectedGraph.SetEdge(edge)
		edgesList.Next()
		fmt.Println("added : ", edge)
		edge = edgesList.Edge()
	}

	connectedNodes := topo.ConnectedComponents(simpleGraph.UndirectedGraph)
	//TODO -- Debug here.
	return connectedNodes
}
