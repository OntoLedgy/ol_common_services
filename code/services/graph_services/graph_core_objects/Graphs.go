package graph_core_objects

import (
	"fmt"
	gonumgraph "gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"sort"
)

type Graphs struct {
	GraphName string
	GraphType string

	*simple.DirectedGraph
	//nodeAdder  gonumgraph.NodeAdder
	graphNodes []*Nodes
	graphEdges []*Edges
}

func (graph *Graphs) AddNodes(
	nodeUuid string,
	NodeDisplayName string) *Nodes {

	newNode := CreateNode(
		nodeUuid,
		NodeDisplayName,
		graph)

	prev := len(
		gonumgraph.NodesOf(
			graph.Nodes()))

	graph.AddNode(newNode.Node)

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

func (graph *Graphs) AddEdges(
	edgeUuid string,
	edgeDisaplyName string,
	node1 *Nodes,
	node2 *Nodes) *Edges {

	newEdge := CreateEdge(
		edgeUuid,
		edgeDisaplyName,
		node1,
		node2,
		graph)

	graph.SetEdge(newEdge.Edge)

	graph.graphEdges =
		append(
			graph.graphEdges,
			newEdge)

	return newEdge
}

//-------------TO BE REVIEWED - REFACTORED, COPIED FROM GONUM------//

//func (graph *Graphs) AddNodeSet(
//	nodeUuid string,
//	NodeDisplayName string) gonumgraph.Node {
//
//	n := 10
//
//	defer func() {
//		r := recover()
//		if r != nil {
//			fmt.Printf("unexpected panic: %v", r)
//			panic(r)
//		}
//	}()
//
//	var addedNodes []gonumgraph.Node
//
//	for i := 0; i < n; i++ {
//		node := graph.NewNode()
//
//		prev := len(gonumgraph.NodesOf(graph.Nodes()))
//		if graph.Node(node.ID()) != nil {
//			curr := graph.Nodes().Len()
//			if curr != prev {
//				fmt.Printf("NewNode mutated graph: prev graph order != curr graph order, %d != %d", prev, curr)
//			}
//			fmt.Printf("NewNode returned existing: %#v", node)
//		}
//		graph.AddNode(node)
//		addedNodes = append(addedNodes, node)
//		curr := len(gonumgraph.NodesOf(graph.Nodes()))
//		if curr != prev+1 {
//			fmt.Printf("AddNode failed to mutate graph: curr graph order != prev graph order+1, %d != %d", curr, prev+1)
//		}
//		if graph.Node(node.ID()) == nil {
//			fmt.Printf("AddNode failed to add node to graph trying to add %#v", node)
//		}
//	}
//
//	ByID(addedNodes)
//
//	graphNodes := gonumgraph.NodesOf(graph.Nodes())
//
//	ByID(graphNodes)
//	if !reflect.DeepEqual(addedNodes, graphNodes) {
//		if n > 20 {
//			fmt.Printf("unexpected node set after node addition: got len:%v want len:%v", len(graphNodes), len(addedNodes))
//		} else {
//			fmt.Printf("unexpected node set after node addition: got:\n %v\nwant:\n%v", graphNodes, addedNodes)
//		}
//	}
//
//	it := graph.Nodes()
//
//	for it.Next() {
//		panicked := panics(func() {
//			graph.AddNode(it.Node())
//		})
//		if !panicked {
//			fmt.Printf("expected panic adding existing node: %v", it.Node())
//		}
//	}
//
//	if gwi, ok := graph.nodeAdder.(gonumgraph.NodeWithIDer); ok {
//		// Test existing nodes.
//		it := graph.Nodes()
//		for it.Next() {
//			id := it.Node().ID()
//			n, new := gwi.NodeWithID(id)
//			if n == nil {
//				fmt.Printf("unexpected nil node for existing node with ID=%d", id)
//			}
//			if new {
//				fmt.Printf("unexpected new node for existing node with ID=%d", id)
//			}
//		}
//		// Run n rounds of ID-specified node addition.
//		for i := 0; i < n; i++ {
//			id := graph.NewNode().ID() // Get a guaranteed non-existing node.
//			n, new := gwi.NodeWithID(id)
//			if n == nil {
//				// Could not create a node, valid behaviour.
//				continue
//			}
//			if !new {
//				fmt.Printf("unexpected old node for non-existing node with ID=%d", id)
//			}
//			graph.AddNode(n) // Use the node to advance to a new non-existing node.
//		}
//	}
//
//	return graph.Node(1)
//}

func ByID(n []gonumgraph.Node) {
	sort.Slice(n, func(i, j int) bool { return n[i].ID() < n[j].ID() })
}

func panics(fn func()) (ok bool) {
	defer func() {
		ok = recover() != nil
	}()
	fn()
	return
}
