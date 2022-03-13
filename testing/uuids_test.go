package testing

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service/uuid_helpers"
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

func TestUuidFactory(t *testing.T) {

	newuuidbase85 := uuid_helpers.CreateNewUuidString()

	fmt.Println("uuid base 85:", newuuidbase85)

	fmt.Println("uuid string:", uuid_helpers.CreateNewUuid())

	fmt.Println("uuid from base 85:", uuid_helpers.CreateUuidFromBase85String(newuuidbase85).UUID.String())
}
