package tree_services

import "github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services/node_lists"
import "github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services/string_processing"

//TODO - move pacakges to internal
//TODO - add other input formats
//TODO - add other output formats

func CreateTrie(string_list_to_process []string) *node_lists.NodeLists {

	sorted_string_list :=
		string_processing.
			Sort_string_list(
				string_list_to_process)

	node_list :=
		node_lists.Create(
			sorted_string_list)

	node_list.
		ProcessStringList(
			sorted_string_list)

	return node_list
}
