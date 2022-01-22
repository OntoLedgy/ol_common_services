package testing

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
	"testing"
)

func TestUuidsTesting(t *testing.T) {

	uuid, err := identity_management_services.GetUUID(
		1,
		"")

	if err == nil {
		fmt.Println(uuid)

	} else {
		fmt.Println(uuid)
	}

}
