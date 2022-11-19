package testing

import (
	"encoding/csv"
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/graph_core_objects"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/path_services"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/same_as_processing"
	yourbasic "github.com/yourbasic/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"log"
	"math"
	"os"
	"strconv"
	"testing"
)

func TestGraph(t *testing.T) {

	filePath := `D:\\S\\go\\src\\github.com\\OntoLedgy\\ol_common_services\\testing\\data\\same_as_links.csv`

	records := getGraphData(filePath)
	fmt.Println(records)

	simpleGraph := graph_core_objects.CreateNewGraph("directed")

	for index, record := range records {

		fmt.Printf("index %v, record : %s\n", index, record)

		simpleGraph.AddNode(
			record[0])

	}
	fmt.Printf("graph %s\n", simpleGraph.DirectedGraph.Nodes())
}

func TestGraphYourBasic(t *testing.T) {

	filePath := `D:\\S\\go\\src\\github.com\\OntoLedgy\\ol_common_services\\testing\\data\\same_as_links.csv`

	records := getGraphData(filePath)
	fmt.Println(records)

	simpleGraph := yourbasic.New(20)

	for index, record := range records {

		fmt.Printf("index %v, record : %s\n", index, record)
		place1, _ := strconv.Atoi(record[0])
		place2, _ := strconv.Atoi(record[1])

		simpleGraph.Add(
			place1,
			place2)
	}
	fmt.Printf("graph:%v", simpleGraph)
	fmt.Printf("graph is Acyclic? : %v\n", yourbasic.Acyclic(simpleGraph))
	fmt.Printf("graph components: %v\n", yourbasic.Components(simpleGraph))

}

func TestSameAsProcessing(t *testing.T) {

	filePath := `D:\\S\\go\\src\\github.com\\OntoLedgy\\ol_common_services\\testing\\data\\same_as_links.csv`

	records := getGraphData(filePath)

	simpleGraph := graph_core_objects.CreateNewGraph("undirected")

	for index, record := range records {

		fmt.Printf("index %v, record : %s\n", index, record)

		node1 := simpleGraph.AddNode(
			record[0])

		node2 := simpleGraph.AddNode(
			record[1])

		simpleGraph.AddEdge("", node1, node2)

	}
	fmt.Printf("graph %s\n", simpleGraph.Nodes())

	sameAsEdges := simpleGraph.Edges()
	sameAsEdges.Reset()
	sameAsEdges.Next()
	sameAsEdge := sameAsEdges.Edge()

	fmt.Printf("first edge: %s", sameAsEdge)

	var sameAsNodeGroups = same_as_processing.ProcessSameAsLinks(sameAsEdges)

	fmt.Printf("same as node group 1 : %s", sameAsNodeGroups[0])
}

func TestNodes(t *testing.T) {

	simpleGraph := graph_core_objects.CreateNewGraph("directed")

	newNode1 := simpleGraph.AddNode(
		"testing 12345")

	newNode2 := simpleGraph.AddNode(
		"testing 12345")

	fmt.Printf("new node 1:%v\nnew node 2: %v\n", newNode1, newNode2)
	fmt.Printf("graph %s\n", simpleGraph.DirectedGraph.Nodes())
}

func TestEdges(t *testing.T) {

	simpleGraph := graph_core_objects.CreateNewGraph("directed")

	newNode1 := simpleGraph.AddNode(
		"testing 123458")

	newNode2 := simpleGraph.AddNode(
		"testing 12345")

	newEdge := simpleGraph.AddEdge(
		"edge uuids",
		newNode1,
		newNode2)

	fmt.Printf("new node 1:%v\nnew node 2: %v\n", newNode1, newNode2)
	fmt.Printf("new edge 1:%v\n", newEdge)
	fmt.Printf("graph %s\n", simpleGraph.DirectedGraph.Nodes())
}

func TestPath(t *testing.T) {

	edges := []simple.WeightedEdge{
		{F: simple.Node('A'), T: simple.Node('B'), W: 9},
		{F: simple.Node('B'), T: simple.Node('C'), W: 14},
		{F: simple.Node('C'), T: simple.Node('D'), W: 15},
		{F: simple.Node('D'), T: simple.Node('E'), W: 12},
		{F: simple.Node('E'), T: simple.Node('F'), W: 10},
		{F: simple.Node('E'), T: simple.Node('B'), W: 8},
		{F: simple.Node('B'), T: simple.Node('F'), W: 7},
		{F: simple.Node('A'), T: simple.Node('E'), W: 8},
		{F: simple.Node('A'), T: simple.Node('F'), W: 14},
		{F: simple.Node('F'), T: simple.Node('D'), W: 12},
		{F: simple.Node('C'), T: simple.Node('F'), W: 8},
	}

	testWeightedDirectedGraph := simple.NewWeightedUndirectedGraph(0, math.Inf(1))

	spanningTreeWeightedDirectedGraph := simple.NewWeightedUndirectedGraph(0, math.Inf(1))

	for _, e := range edges {
		testWeightedDirectedGraph.SetWeightedEdge(e)
	}

	fmt.Println("checking Minimum Spanning Tree using Prim")

	minimumSpanningTreePrimValue := path.Prim(spanningTreeWeightedDirectedGraph, testWeightedDirectedGraph)

	fmt.Printf("minimum spanning Tree Prim value: %f\n", minimumSpanningTreePrimValue)

	fmt.Println("getting Minimum Spanning Tree using path and cost")

	path_services.GetMinimumSpanningTree(testWeightedDirectedGraph)

}

func getGraphData(filePath string) [][]string {

	file, fileOpenError := os.Open(filePath)

	if fileOpenError != nil {
		panic("cant open file")
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	records, fileReadError := csvReader.ReadAll()

	if fileReadError != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, fileReadError)
	}

	return records
}
