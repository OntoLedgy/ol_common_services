package types

//from enum import auto
//from enum import unique
//
//from nf_common_source.code.nf.types.column_types import ColumnTypes

//@unique
//class NfColumnTypes(
type OlColumnTypes int

const (
	//ColumnTypes):
	//NF_UUIDS = \
	//auto()
	OL_UUIDS OlColumnTypes = iota
	//UNIVERSE_KEYS = \
	UNIVERSE_KEYS
	//auto()
	COLLECTION_TYPES
	//COLLECTION_TYPES = \
	//auto()
)

//column_name = \
//property(
//fget=__column_name)

//def __column_name(
//self) \
//-> str:
func (olColumnTypes OlColumnTypes) ColumnName() string {

	//column_name = \
	//column_name_mapping[self]

	//column_name_mapping = \
	//{

	//}
	switch olColumnTypes {

	//NfColumnTypes.NF_UUIDS: 'nf_uuids',
	case OL_UUIDS:
		return "nf_uuids"
	//NfColumnTypes.UNIVERSE_KEYS: 'universe_keys',
	case UNIVERSE_KEYS:
		return "universe_keys"
	//NfColumnTypes.COLLECTION_TYPES: 'collection_types'
	case COLLECTION_TYPES:
		return "collection_types"

	}

	//return \
	//column_name

	return "unknown "
}
