package types

//https://github.com/boro-alpha/nf_common/blob/master/nf_common_source/code/nf/types/collection_types.py
//from enum import Enum

//class CollectionTypes(
//Enum):
type CollectionTypes int

//TODO : Needs to be  implemented
//def __collection_name(
//self) \
//-> str:
//raise \
//NotImplementedError
//
//collection_name = \
//property(
//fget=__collection_name)
func (collectionType CollectionTypes) CollectionName() string {

	return "no implemented"
}
