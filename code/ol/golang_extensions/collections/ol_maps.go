package collections

//from nf_common_source.code.nf.python_extensions.collections.nf_dictionary_try_get_results import \
//NfDictionaryTryGetResults

//class NfDictionaries(
//dict):
type OlMaps map[any]any

//def __init__(
//self):

func (olMap OlMaps) Initialise() {
	//dict.__init__(
	//self)
	olMap = map[any]any{}

	return
}

func (olMap OlMaps) Keys() []any {

	keys := make([]any, len(olMap))
	for key, _ := range olMap {
		keys = append(keys, key)
	}
	return keys
}

//def try_get_value(
//self,
//key) \
//-> NfDictionaryTryGetResults:
func (olMap OlMaps) TryGetValue(key any) OlMapTryGetResults {

	//nf_dictionary_try_get_result = \
	//NfDictionaryTryGetResults()
	olMapTryGetResult := OlMapTryGetResults{}

	//nf_dictionary_try_get_result.value = \
	//self.get(
	//key)

	//nf_dictionary_try_get_result.key_exists = \
	//key in self.keys()
	value, keyExists := olMap[key]

	//if nf_dictionary_try_get_result.key_exists:
	if keyExists {
		olMapTryGetResult.SetKeyExists(true)
		olMapTryGetResult.SetValue(value)
		//else:
	} else {
		//nf_dictionary_try_get_result.value = \
		//None
		olMapTryGetResult.SetValue(nil)

	}

	//return \
	//nf_dictionary_try_get_result
	return olMapTryGetResult
}
