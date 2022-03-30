package trees

import (
	"fmt"
)

type Trees struct {
	RootNodeValue any
	children      []*Trees
}

func (tree *Trees) AddChild(
	nodeToAdd *Trees) {

	tree.children =
		append(
			tree.children,
			nodeToAdd)

}

func (tree *Trees) RemoveChild(
	nodeToRemove *Trees) {

	for index, childNode := range tree.children {
		if nodeToRemove == childNode {
			tree.children =
				append(
					tree.children[:index],
					tree.children[index+1:]...)
		}
	}

}

func (tree *Trees) CheckIfNodeIsAChild(
	childNode *Trees) bool {

	for _, currentNodeChildNode := range tree.children {

		if childNode == currentNodeChildNode {
			return true
		}

	}
	return false
}

func (tree *Trees) PrintTree() {

	fmt.Printf("root node: %s\n", tree.RootNodeValue)

	if tree.children != nil {
		fmt.Printf("	children of %s:\n", tree.RootNodeValue)

		for _, child := range tree.children {

			child.PrintTree()

		}
		fmt.Printf("\n	end of children for %s.\n", tree.RootNodeValue)
	} else {

		fmt.Printf("end of %s\n", tree.RootNodeValue)
	}

}
