package uuid_service

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type UUIDs struct {
	uuid.UUID
}

func GetUUID(
	uuidType int,
	seedingString ...string) (
	*UUIDs,
	error) {

	var uuid_error error

	generated_uuid := &UUIDs{}

	generated_uuid.UUID = uuid.UUID{}

	switch uuidType {

	case 1:
		// Creating UUIDs Version 4
		generated_uuid.UUID =
			uuid.NewV4()

	case 2:
		// Parsing UUIDs from string input
		generated_uuid.UUID, uuid_error =
			uuid.FromString(
				seedingString[0])

		if uuid_error != nil {
			fmt.Printf("Something went wrong: %s", uuid_error)
			return nil, uuid_error
		}

	default:

	}
	return generated_uuid, nil

}
