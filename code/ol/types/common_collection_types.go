package types

//https://github.com/boro-alpha/nf_common/blob/master/nf_common_source/code/nf/types/common_collection_types.py

//from enum import auto
//from enum import unique
//from nf_common_source.code.nf.types.collection_types import CollectionTypes

//@unique
//class CommonCollectionTypes(
//CollectionTypes):
type CommonCollectionTypes int

const (
	//SUMMARY_TABLE = auto()
	SUMMARY_TABLE CommonCollectionTypes = iota
)

//def __collection_name(
//self) \
//-> str:
//collection_name = \
//collection_name_mapping[self]
//
//return \
//collection_name
//collection_name = \
//property(
//fget=__collection_name)

func (commonCollectionType CommonCollectionTypes) CollectionName() string {

	//collection_name_mapping = \
	switch commonCollectionType {
	//CommonCollectionTypes.SUMMARY_TABLE: 'summary_table'
	case SUMMARY_TABLE:
		return "summary_table"
	}
	return "not implemented"
}
