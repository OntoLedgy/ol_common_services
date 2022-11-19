package path_services

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

//TODO - add infrastructure to select algorithm - currently defaulting to FloydWarsh

func GetShortestPaths(testWeightedDirectedGraph graph.Graph) path.AllShortest {

	//shortestPaths := path.DijkstraAllPaths(*testWeightedDirectedGraph)

	// Find the shortest path to each node from Q.
	shortestPaths, ok := path.FloydWarshall(testWeightedDirectedGraph)

	if ok {
		fmt.Println("no negative cycle present")

	}

	return shortestPaths
}

func GetMinimumSpanningTree(testWeightedDirectedGraph *simple.WeightedUndirectedGraph) {

	allNodePermutationPaths := getAllPermutationsOfNodes(testWeightedDirectedGraph)

	shortestPaths := path.DijkstraAllPaths(testWeightedDirectedGraph)

	nodePathCostSet := make(map[int]float64)

	for nodePermutationIndex, nodePermutation := range allNodePermutationPaths {

		permutationCost := getPathCost(nodePermutation, shortestPaths)

		nodePathCostSet[nodePermutationIndex] = permutationCost
	}
	minimumSpanningTreeCost, minimumSpanningTreeIndicies := getSmallestValues(nodePathCostSet)
	fmt.Printf("shortest path length:%v\n", minimumSpanningTreeCost)
	for _, minimumSpanningTreeIndex := range minimumSpanningTreeIndicies {

		fmt.Printf("shortest path :%c\n", allNodePermutationPaths[minimumSpanningTreeIndex])

	}

}

func getAllPermutationsOfNodes(currentGraph *simple.WeightedUndirectedGraph) [][]graph.Node {

	//var path [][]graph.Node

	var nodeSet []graph.Node

	nodes := currentGraph.Nodes()
	nodes.Reset()
	nodes.Next()
	numberOfNodes := nodes.Len()

	for index := 0; index <= numberOfNodes; nodes.Next() {
		node := nodes.Node()
		nodeSet = append(nodeSet, node)

		index++
	}

	allNodePermutationPaths := permutations(nodeSet)

	return allNodePermutationPaths

}

func getSmallestValues(nodePathCostSet map[int]float64) (float64, []int) {

	shortestPathCost := 10000000000.00

	for _, nodePathCost := range nodePathCostSet {
		if nodePathCost < shortestPathCost {
			shortestPathCost = nodePathCost
		}
	}

	var shortestValueIndices []int

	for index, nodePathCost := range nodePathCostSet {
		if nodePathCost == shortestPathCost {
			shortestValueIndices = append(shortestValueIndices, index)
		}

	}

	return shortestPathCost, shortestValueIndices
}

func getPathCost(nodePermutation []graph.Node, allShortestPaths path.AllShortest) float64 {

	pathCost := 0.00

	currentNode := nodePermutation[0]

	//fmt.Printf("calculating node permutation : %c\n", nodePermutation)

	for index := 0; index <= len(nodePermutation)-1; index++ {

		pathCost += allShortestPaths.Weight(currentNode.ID(), nodePermutation[index].ID())
		currentNode = nodePermutation[index]
		//fmt.Printf("cost : %v\n", pathCost)
	}

	if pathCost == 44 {
		fmt.Printf("calculating node permutation : %c\n", nodePermutation)
		fmt.Printf("lowest total cost--------------------------------: %v\n", pathCost)
	}

	return pathCost
}

func permutations(arr []graph.Node) [][]graph.Node {

	var helper func([]graph.Node, int)

	res := [][]graph.Node{}

	helper = func(arr []graph.Node, n int) {

		if n == 1 {
			tmp := make([]graph.Node, len(arr))

			copy(tmp, arr)

			res = append(res, tmp)

		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))

	return res
}
