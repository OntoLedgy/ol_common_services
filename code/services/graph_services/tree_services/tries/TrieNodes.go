package tries

import (
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services/trees"
)

type Tries struct {
	Nodes []*trees.Trees
}

func (trie *Tries) PopulateList(listOfStrings []string) {

	for _, string := range listOfStrings {

		trie.
			addNode(
				string)

	}
}

func (trie *Tries) addNode(stringToAdd string) {

	node :=
		trees.
			Create(stringToAdd)

	trie.
		Nodes =
		append(trie.Nodes, node)
}

func (trie *Tries) ProcessStringList(
	sortedStringList []string) {

	for _, stringInList := range sortedStringList {

		trie.processStringInList(
			stringInList,
			sortedStringList)
	}

}

func (trie *Tries) processStringInList(
	stringInList string,
	sortedStringList []string) {

	var node *trees.Trees

	node =
		trie.
			GetNode(
				stringInList)

	descendants :=
		FindAllDescendants(
			stringInList,
			sortedStringList)

	trie.
		addDescendants(
			descendants)

	trie.
		updateChildrenForANode(
			descendants,
			node)

}

func (trie *Tries) updateChildrenForANode(
	childrenStrings []string,
	node *trees.Trees) {

	for _, childString := range childrenStrings {

		child_node :=
			trie.
				GetNode(childString)

		node.
			AddChild(child_node)

		trie.
			RemoveChildFromOtherNodes(
				node,
				child_node)

	}

}

func (trie *Tries) GetNode(
	stringToCheck string) *trees.Trees {

	var node *trees.Trees

	if trie.findNode(stringToCheck) == nil {

		node = trees.Create(stringToCheck)
		trie.Nodes = append(trie.Nodes, node)

	} else {
		node = trie.findNode(stringToCheck)
	}

	return node
}

func (trie *Tries) findNode(
	stringValue string) *trees.Trees {

	for _, node := range trie.Nodes {

		if node.RootNodeValue == stringValue {

			return node

		}

	}
	return nil
}

func (trie *Tries) RemoveChildFromOtherNodes(
	node *trees.Trees,
	childNode *trees.Trees) {

	for _, current_node := range trie.Nodes {

		if current_node != node {

			if current_node.CheckIfNodeIsAChild(childNode) {

				current_node.RemoveChild(childNode)
			}

		}

	}

}

func (trie *Tries) addDescendants(
	descendants []string) {

	for _, descendant := range descendants {

		trie.
			GetNode(
				descendant)

	}
}
