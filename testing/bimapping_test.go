package testing

import (
	"github.com/OntoLedgy/ol_common_services/code/ol/golang_extensions/collections"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services/trees"
	"github.com/OntoLedgy/ol_common_services/code/services/graph_services/tree_services/tries"
	"testing"
)

func TestBimapping(t *testing.T) {

	object1 := new(trees.Trees)

	object2 := new(tries.Tries)

	bimapping := new(collections.OlBimappings)

	bimapping.Initialise(nil)

	bimapping.AddMapping(object1, object2)

	bimapping.GetDomainKeyedOnRange()

}
