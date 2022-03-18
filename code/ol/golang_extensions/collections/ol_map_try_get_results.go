package collections

//class NfDictionaryTryGetResults:
type OlMapTryGetResults struct {
	keyExits bool
	Key      any
	value    any
}

//def __init__(
//self):

func (olMapTryGetResult OlMapTryGetResults) Initalise() {

	//self.__key_exists = \
	olMapTryGetResult.keyExits =
		//False
		false

	//self.__value = \
	olMapTryGetResult.value =
		//None
		nil
}

//def __get_key_exists(
//self) \
//-> bool:
func (olMapTryGetResult OlMapTryGetResults) KeyExists() bool {

	//key_exists = \
	keyExists :=
		//self.__key_exists
		olMapTryGetResult.keyExits

	//return \
	//key_exists
	return keyExists

}

//def __set_key_exists(
//self,
func (olMapTryGetResult OlMapTryGetResults) SetKeyExists(
	//key_exists: bool):
	keyExists bool) {

	//self.__key_exists = \
	//key_exists
	olMapTryGetResult.keyExits = keyExists
}

//def __get_value(
func (olMapTryGetResult OlMapTryGetResults) Value() any {
	//self):
	value :=
		olMapTryGetResult.value
	//value = \
	//self.__value
	//
	//return \
	//value
	return value
}

//def __set_value(
func (olMapTryGetResult OlMapTryGetResults) SetValue(
	value any) {
	//self,
	//value):

	//self.__value = \
	olMapTryGetResult.value =
		//value
		value

}

//key_exists = \
//property(
//fget=__get_key_exists,
//fset=__set_key_exists)

//value = \
//property(
//fget=__get_value,
//fset=__set_value)
