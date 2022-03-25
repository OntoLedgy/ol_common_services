package nodes

import (
	"fmt"
)

type TreeNodes struct {
	Value    string
	children []*TreeNodes
}

func (node *TreeNodes) AddChild(
	node_to_add *TreeNodes) {

	node.children =
		append(
			node.children,
			node_to_add)

}

func (node *TreeNodes) RemoveChild(
	node_to_remove *TreeNodes) {

	for index, child_node := range node.children {
		if node_to_remove == child_node {
			node.children = append(node.children[:index], node.children[index+1:]...)
		}
	}

}

func (current_node *TreeNodes) CheckIfNodeIsAChild(
	child_node *TreeNodes) bool {

	for _, current_node_child_node := range current_node.children {

		if child_node == current_node_child_node {
			return true
		}

	}
	return false
}

func (node *TreeNodes) PrintNode() {

	fmt.Printf("node: %s\n", node.Value)

	if node.children != nil {
		fmt.Printf("	children of %s:\n", node.Value)

		for _, child := range node.children {

			child.PrintNode()

		}
		fmt.Printf("\n	end of children for %s.\n", node.Value)
	} else {

		fmt.Printf("end of %s\n", node.Value)
	}

}
