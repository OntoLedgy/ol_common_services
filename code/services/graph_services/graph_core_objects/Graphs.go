package graph_core_objects

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service/uuid_helpers"
	gonumgraph "gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type Graphs struct {
	GraphName string
	GraphType string

	*simple.DirectedGraph
	*simple.UndirectedGraph

	graphNodes []*Nodes
	graphEdges []*Edges
}

func (graph *Graphs) AddNode(
	NodeDisplayName string) *Nodes {

	nodeUuid := uuid_helpers.CreateNewUuid4()

	newNode := CreateNode(
		nodeUuid,
		NodeDisplayName,
		graph)

	prev := len(
		gonumgraph.NodesOf(
			graph.Nodes()))

	graph.addNode(newNode.Node)

	curr := len(
		gonumgraph.NodesOf(
			graph.Nodes()))

	if curr != prev+1 {
		fmt.Printf("AddNode failed to mutate graph: curr graph order != prev graph order+1, %d != %d", curr, prev+1)
	}

	if graph.Node(newNode.Node.ID()) == nil {
		fmt.Printf("AddNode failed to add node to graph trying to add %#v", newNode)
	}

	graph.graphNodes = append(
		graph.graphNodes,
		newNode)

	return newNode
}

func (graph *Graphs) addNode(
	node gonumgraph.Node) {

	switch graph.GraphType {
	case "directed":
		graph.DirectedGraph.AddNode(node)
	case "undirected":
		graph.UndirectedGraph.AddNode(node)
	default:
		graph.DirectedGraph.AddNode(node)
	}

}

func (graph *Graphs) NewNode() gonumgraph.Node {

	switch graph.GraphType {
	case "directed":
		return graph.DirectedGraph.NewNode()
	case "undirected":
		return graph.UndirectedGraph.NewNode()
	default:
		return graph.DirectedGraph.NewNode()
	}

}

func (graph *Graphs) Nodes() gonumgraph.Nodes {

	switch graph.GraphType {
	case "directed":
		return graph.DirectedGraph.Nodes()
	case "undirected":
		return graph.UndirectedGraph.Nodes()
	default:
		return graph.DirectedGraph.Nodes()

	}
}

func (graph *Graphs) Node(
	nodeId int64) gonumgraph.Node {
	switch graph.GraphType {
	case "directed":
		return graph.DirectedGraph.Node(nodeId)
	case "undirected":
		return graph.UndirectedGraph.Node(nodeId)
	default:
		return graph.DirectedGraph.Node(nodeId)

	}
}

func (graph *Graphs) AddEdge(
	edgeDisplayName string,
	node1 *Nodes,
	node2 *Nodes) *Edges {

	edgeUuid := uuid_helpers.CreateNewUuid4()
	newEdge := CreateEdge(
		edgeUuid,
		edgeDisplayName,
		node1,
		node2,
		graph)

	prev := len(
		gonumgraph.EdgesOf(
			graph.Edges()))

	graph.addEdge(newEdge)

	curr := len(
		gonumgraph.EdgesOf(
			graph.Edges()))

	if curr != prev+1 {
		fmt.Printf("AddNode failed to mutate graph: curr graph order != prev graph order+1, %d != %d", curr, prev+1)
	}

	if graph.Edge(node1.ID(), node2.ID()) == nil {
		fmt.Printf("AddEdge failed to add node to graph trying to add %#v", newEdge)
	}

	graph.graphEdges = append(
		graph.graphEdges,
		newEdge)

	return newEdge

}

func (graph *Graphs) addEdge(
	edge gonumgraph.Edge) {

	switch graph.GraphType {

	case "directed":
		graph.DirectedGraph.SetEdge(edge)
	case "undirected":
		graph.UndirectedGraph.SetEdge(edge)
	default:
		graph.DirectedGraph.SetEdge(edge)
	}

}

func (graph *Graphs) Edges() gonumgraph.Edges {
	switch graph.GraphType {
	case "directed":
		return graph.DirectedGraph.Edges()
	case "undirected":
		return graph.UndirectedGraph.Edges()
	default:
		return graph.DirectedGraph.Edges()

	}
}

func (graph *Graphs) Edge(
	node1Id,
	node2Id int64) gonumgraph.Edge {

	switch graph.GraphType {
	case "directed":
		return graph.DirectedGraph.Edge(node1Id, node2Id)
	case "undirected":
		return graph.UndirectedGraph.Edge(node1Id, node2Id)
	default:
		return graph.DirectedGraph.Edge(node1Id, node2Id)

	}

}
