package testing

import (
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services/tries"
	"testing"
)

func TestDataTrie(t *testing.T) {

	string_list_to_process := []string{
		"123",
		"1231",
		"12312",
		"123124"}

	trie := tries.Create(string_list_to_process)

	trie.Nodes[0].PrintTree()

}
