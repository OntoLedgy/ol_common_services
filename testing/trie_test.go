package testing

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services"
	"testing"
)

func TestDataTrie(t *testing.T) {

	string_list_to_process := []string{
		"123213",
		"12321",
		"12321",
		"123455435",
		"12324",
		"12234567"}

	trie := tree_services.CreateTrie(string_list_to_process)

	fmt.Println(trie.Nodes)

}

//TODO - move test data out to testing
