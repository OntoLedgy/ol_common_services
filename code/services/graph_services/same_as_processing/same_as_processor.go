package same_as_processing

import "github.com/OntoLedgy/ol_common_services/code/services/graph_services/graph_core_objects"

func process_same_as_list(edgesList []*graph_core_objects.Edges) []*graph_core_objects.Edges {

	simpleGraph := graph_core_objects.CreateNewGraph(0)

	simpleGraph.DirectedGraph.Edges()

	return edgesList
}
