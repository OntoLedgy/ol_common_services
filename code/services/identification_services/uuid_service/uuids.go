package uuid_service

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type UUIDs struct {
	*uuid.UUID
}

func GetUUID(
	uuid_type int,
	seeding_string string) (
	*UUIDs,
	error) {

	var uuid_error error

	generated_uuid := new(UUIDs)

	generated_uuid.UUID = new(uuid.UUID)

	switch uuid_type {

	case 1:
		// Creating UUIDs Version 4
		*generated_uuid.UUID =
			uuid.NewV4()

	case 2:
		// Parsing UUIDs from string input
		*generated_uuid.UUID, uuid_error =
			uuid.FromString(seeding_string)

		if uuid_error != nil {
			fmt.Printf("Something went wrong: %s", uuid_error)
			return nil, uuid_error
		}

	}
	return generated_uuid, nil

}
