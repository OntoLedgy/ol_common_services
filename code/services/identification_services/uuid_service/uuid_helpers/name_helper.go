package uuid_helpers

import "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service/constants"

//def get_uuidified_dataframe_name(
//dataframe_name: str):
func GetUuidifiedDataframeName(
	dataFrameName string) string {

	//uuidified_dataframe_name = \
	uuidifiedDataFrameName :=
		//__get_separated_uuid_column(
		getSeperatedUuidColumn(
			//UUIDIFIED_DATAFRAME_PREFIX,
			constants.UUIDIFIED_DATAFRAME_PREFIX,
			//dataframe_name)
			dataFrameName)

	//return \
	//uuidified_dataframe_name

	return uuidifiedDataFrameName

}

//def get_common_uuidification_table_name(
func GetCommonUuidificationTableName(
	//prefix: str,
	prefix string,
	//table_type_name: str):
	tableTypeName string) string {
	//common_uuidification_table_name = \
	commonUuidificationTableName :=
		//__get_separated_uuid_column(
		getSeperatedUuidColumn(
			//prefix,
			prefix,
			//table_type_name)
			tableTypeName)
	//return \
	//common_uuidification_table_name
	return commonUuidificationTableName
}

//def __get_separated_uuid_column(
func getSeperatedUuidColumn(
	//prefix: str,
	prefix string,
	//suffix: str):
	suffix string) string {

	//separated_uuid_column = \
	sepeatedUuidColumn :=
		//prefix + UUID_COLUMN_NAME_SEPARATOR + suffix
		prefix + constants.UUID_COLUMN_NAME_SEPARATOR + suffix

	//return \
	//separated_uuid_column
	return sepeatedUuidColumn

}
