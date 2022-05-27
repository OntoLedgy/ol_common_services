package testing

import (
	"encoding/csv"
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/graph_core_objects"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service/uuid_helpers"
	yourbasic "github.com/yourbasic/graph"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestGraph(t *testing.T) {

	filePath := `D:\\S\\go\\src\\github.com\\OntoLedgy\\ol_common_services\\testing\\data\\same_as_links.csv`

	records := GetGraphData(filePath)
	fmt.Println(records)

	simpleGraph := graph_core_objects.CreateNewGraph(0)

	for index, record := range records {

		fmt.Printf("index %v, record : %s\n", index, record)

		nodeUuid := uuid_helpers.CreateNewUuid4()

		simpleGraph.AddNodes(
			nodeUuid,
			record[0])

	}
	fmt.Printf("graph %s\n", simpleGraph.Nodes())
}

func TestGraphYourBasic(t *testing.T) {

	filePath := `D:\\S\\go\\src\\github.com\\OntoLedgy\\ol_common_services\\testing\\data\\same_as_links.csv`

	records := GetGraphData(filePath)
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
	fmt.Printf("graph is Acyclic ? : %v\n", yourbasic.Acyclic(simpleGraph))
	fmt.Printf("graph components ? : %v\n", yourbasic.Components(simpleGraph))
}

func TestNodes(t *testing.T) {

	simpleGraph := graph_core_objects.CreateNewGraph(0)

	nodeUuid := uuid_helpers.CreateNewUuid4()

	newNode1 := simpleGraph.AddNodes(
		nodeUuid,
		"testing 12345")

	newNode2 := simpleGraph.AddNodes(
		nodeUuid,
		"testing 12345")

	fmt.Printf("new node 1:%v\nnew node 2: %v\n", newNode1, newNode2)
	fmt.Printf("graph %s\n", simpleGraph.Nodes())
}

func TestEdges(t *testing.T) {

	simpleGraph := graph_core_objects.CreateNewGraph(0)

	nodeUuid := uuid_helpers.CreateNewUuid4()

	newNode1 := simpleGraph.AddNodes(
		nodeUuid,
		"testing 12345")

	newNode2 := simpleGraph.AddNodes(
		nodeUuid,
		"testing 12345")

	newEdge := simpleGraph.AddEdges(
		"edge uuids",
		"test edge",
		newNode1,
		newNode2)

	fmt.Printf("new node 1:%v\nnew node 2: %v\n", newNode1, newNode2)
	fmt.Printf("new edge 1:%v\n", newEdge)
	fmt.Printf("graph %s\n", simpleGraph.Nodes())
}

func GetGraphData(filePath string) [][]string {

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
