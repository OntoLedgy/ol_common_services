package uuid_helpers

//
//import uuid
//import base64

import (
	"encoding/ascii85"
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
	uuid "github.com/satori/go.uuid"
)

//def create_new_uuid_string():
func CreateNewUuidString() string {

	//new_uuid = \
	newUuid :=
		//uuid.uuid1()
		uuid.NewV1()

	//uuid_as_base85 = \
	uuidAsBase85Encode := make(
		[]byte,
		ascii85.MaxEncodedLen(len(newUuid)))
	//base64.b85encode(
	ascii85.Encode(
		uuidAsBase85Encode,
		//new_uuid.bytes)
		newUuid.Bytes())

	//uuid_as_base85_string = \
	//uuid_as_base85.decode()
	uuidAsBase85Decode := make(
		[]byte,
		ascii85.MaxEncodedLen(len(uuidAsBase85Encode)))

	ascii85.Decode(uuidAsBase85Decode,
		uuidAsBase85Encode,
		true)

	uuidAsBase85String := string(
		uuidAsBase85Decode)

	//return \
	//uuid_as_base85_string
	return uuidAsBase85String

}

//def create_new_uuid() \
//-> str:
func CreateNewUuid() string {

	//new_uuid = \
	newUuid :=
		//str(uuid.uuid1())
		uuid.NewV1().String()

	//return \
	//new_uuid
	return newUuid

}

func CreateNewUuid4() string {

	//new_uuid = \
	newUuid :=
		//str(uuid.uuid1())
		uuid.NewV4().String()

	//return \
	//new_uuid
	return newUuid

}

//def create_uuid_from_base85_string(
//uuid_as_base85_string):
func CreateUuidFromBase85String(baseAs85String string) uuid_service.UUIDs {

	decodedUuid := make([]byte, 4*len(baseAs85String))

	//uuid_as_bytes = \
	//base64.b85decode(
	//uuid_as_base85_string)
	ascii85.Decode(
		decodedUuid,
		[]byte(baseAs85String),
		true)

	fmt.Println(baseAs85String)
	fmt.Println(decodedUuid)

	//uuid_from_bytes = \
	uuidFromBytes, _ :=
		uuid.FromBytes(
			decodedUuid)
	//uuid.UUID(
	//bytes=uuid_as_bytes)
	uuidFromBytesWrapped :=
		uuid_service.UUIDs{
			uuidFromBytes}

	//return \
	//uuid_from_bytes
	return uuidFromBytesWrapped

}

//def create_uuid_from_canonical_format_string(
//canonical_format_string):
//
//uuid_from_canonical_format = \
//uuid.UUID(
//canonical_format_string)
//
//return \
//uuid_from_canonical_format
