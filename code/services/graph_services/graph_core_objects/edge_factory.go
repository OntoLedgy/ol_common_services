package graph_core_objects

func CreateEdge(
	edgeUuid string,
	edgeDisaplyName string,
	node1 *Nodes,
	node2 *Nodes,
	graph *Graphs) *Edges {

	edge := graph.NewEdge(
		node1,
		node2)

	newEdge := &Edges{
		edgeUuid,
		edgeDisaplyName,
		node1,
		node2,
		edge}

	return newEdge
}
