package testing

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
	"testing"
)

func TestUuidsTesting(t *testing.T) {

	uuid, err := uuid_service.GetUUID(
		1,
		"")

	if err == nil {
		fmt.Println(uuid)

	} else {
		fmt.Println(uuid)
	}

}
