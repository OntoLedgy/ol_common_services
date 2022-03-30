package collections

//https://github.com/boro-alpha/nf_common/blob/master/nf_common_source/code/nf/python_extensions/collections/nf_bimappings.py

//from nf_common_source.code.nf.python_extensions.collections.nf_dictionaries import NfDictionaries

//class NfBimappings:
type OlBimappings struct {
	domainKeyedOnRange *OlMaps
	rangeKeyedOnDomain *OlMaps
}

//def __init__(
//self,
//map: dict):
func (olBimapping OlBimappings) Initialise(mapForBiMapping map[any]any) {

	//NfBimappings.__populate_internal_dictionaries(
	olBimapping.populateInternalDictionaries(
		//map=map)
		mapForBiMapping)
	//)
}

//def try_get_range_using_domain(
//self,
//domain_key):
func (olBimapping OlBimappings) TryGetRangeUsingDomain(domainKey any) any {

	//range_value = \
	rangeValue :=
		//self.__range_keyed_on_domain.try_get_value(
		olBimapping.rangeKeyedOnDomain.TryGetValue(
			//key=domain_key)
			domainKey)

	//return \
	//range_value
	return rangeValue

}

//def try_get_domain_using_range(
//self,
//range_key):
func (olBimapping OlBimappings) TryGetDomainUsingRange(rangeKey any) any {

	//domain_value = \
	domainValue :=
		//self.__domain_keyed_on_range.try_get_value(
		olBimapping.domainKeyedOnRange.TryGetValue(
			//key=rangeKey)
			rangeKey)

	//return \
	//domain_value
	return domainValue
}

//def getRange(
//self):
func (olBimapping OlBimappings) GetRange() []any {

	//range = \
	mapRange :=
		//self.__domain_keyed_on_range.keys()
		olBimapping.domainKeyedOnRange.Keys()

	//return \
	//range
	return mapRange
}

//def get_domain(
func (olBimapping OlBimappings) GetDomain() []any {

	//self):
	//domain = \
	mapDomain :=
		//self.__range_keyed_on_domain.keys()
		olBimapping.rangeKeyedOnDomain.Keys()

	//return \
	//domain
	return mapDomain
}

//def get_range_keyed_on_domain(
//self):
func (olBimapping OlBimappings) GetRangeKeyedOnDomain() *OlMaps {

	//range_keyed_on_domain = \
	rangeKeyedOnDomain :=
		//self.__range_keyed_on_domain
		olBimapping.rangeKeyedOnDomain

	//return \
	//range_keyed_on_domain
	return rangeKeyedOnDomain

}

//def get_domain_keyed_on_range(
//self):
func (olBimapping OlBimappings) GetDomainKeyedOnRange() *OlMaps {

	//domain_keyed_on_range = \
	domainKeyedOnRange :=
		//self.__domain_keyed_on_range
		olBimapping.domainKeyedOnRange
	//return \
	//domain_keyed_on_range
	return domainKeyedOnRange
}

//def __populate_internal_dictionaries(
//self,
//map: dict):
func (olBimapping OlBimappings) populateInternalDictionaries(
	mapForBiMapping map[any]any) {

	//self.__domain_keyed_on_range = \
	//NfDictionaries()
	olBimapping.domainKeyedOnRange =
		new(OlMaps)

	//self.__range_keyed_on_domain = \
	//NfDictionaries()
	olBimapping.rangeKeyedOnDomain =
		new(OlMaps)

	//for domain_value, range_value in map.items():
	for domainValue, rangeValue := range mapForBiMapping {

		//self.add_mapping(
		olBimapping.AddMapping(
			//domain_value=domain_value,
			domainValue,
			//range_value=range_value)
			rangeValue)
	}
}

//def add_mapping(
//self,
func (olBimapping OlBimappings) AddMapping(
	//domain_value,
	DomainValue any,
	//range_value):
	RangeValue any) {

	domainKeyedOnRange := *olBimapping.domainKeyedOnRange
	//self.__domain_keyed_on_range[range_value] = \
	domainKeyedOnRange[RangeValue] =
		//domain_value
		DomainValue

	rangeKeyedOnDomain := *olBimapping.rangeKeyedOnDomain
	rangeKeyedOnDomain[DomainValue] =
		//self.__range_keyed_on_domain[domain_value] = \
		RangeValue
	//range_value
}

//@staticmethod
//def __populate_range_keyed_on_range(
//map: NfDictionaries) -> NfDictionaries:
func (olBimapping OlBimappings) populateRangeKeyedOnRange(
	olMap OlMaps) OlMaps {

	//domain_keyed_on_range = \
	domainKeyedOnRange :=
		//NfDictionaries()
		OlMaps{}

	//for domain_value, range_value in map:
	for domainValue, rangeValue := range olMap {

		//domain_keyed_on_range[range_value] = \
		domainKeyedOnRange[rangeValue] =
			//domain_value
			domainValue
	}

	//return \
	//domain_keyed_on_range
	return domainKeyedOnRange

}
