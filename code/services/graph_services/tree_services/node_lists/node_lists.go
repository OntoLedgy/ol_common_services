package node_lists

import (
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services/nodes"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services/string_processing"
)

type NodeLists struct {
	Nodes []*nodes.TreeNodes
}

func (node_list *NodeLists) PopulateList(list_of_strings []string) {

	for _, string := range list_of_strings {

		node_list.
			addNode(
				string)

	}
}

func (node_list *NodeLists) addNode(string_to_add string) {

	node :=
		nodes.
			Create_node(string_to_add)

	node_list.
		Nodes =
		append(node_list.Nodes, node)
}

func (node_list *NodeLists) ProcessStringList(
	sorted_string_list []string) {

	for _, string_in_list := range sorted_string_list {

		node_list.processStringInList(
			string_in_list,
			sorted_string_list)

	}

}

func (node_list *NodeLists) processStringInList(
	string_in_list string,
	sorted_string_list []string) {

	var node *nodes.TreeNodes

	node =
		node_list.
			GetNode(
				string_in_list)

	decendants :=
		string_processing.
			Find_all_descendants(
				string_in_list,
				sorted_string_list)

	node_list.
		AddDecendants(
			decendants)

	node_list.
		updateChildrenForANode(
			decendants,
			node)

}

func (node_list *NodeLists) updateChildrenForANode(
	children_strings []string,
	node *nodes.TreeNodes) {

	for _, child_string := range children_strings {

		child_node :=
			node_list.
				GetNode(child_string)

		node.
			AddChild(child_node)

		node_list.
			RemoveChildFromOtherNodes(
				node,
				child_node)

	}

}

func (node_list *NodeLists) GetNode(
	string_to_check string) *nodes.TreeNodes {

	var node *nodes.TreeNodes

	if node_list.findNode(string_to_check) == nil {

		node = nodes.Create_node(string_to_check)
		node_list.Nodes = append(node_list.Nodes, node)

	} else {
		node = node_list.findNode(string_to_check)
	}

	return node
}

func (node_list *NodeLists) findNode(
	string_value string) *nodes.TreeNodes {

	for _, node := range node_list.Nodes {

		if node.Value == string_value {

			return node

		}

	}
	return nil
}

func (node_list *NodeLists) RemoveChildFromOtherNodes(
	node *nodes.TreeNodes,
	child_node *nodes.TreeNodes) {

	for _, current_node := range node_list.Nodes {

		if current_node != node {

			if current_node.CheckIfNodeIsAChild(child_node) {

				current_node.RemoveChild(child_node)
			}

		}

	}

}

func (node_list *NodeLists) AddDecendants(
	decendants []string) {

	for _, decendant := range decendants {

		node_list.
			GetNode(
				decendant)

	}
}
