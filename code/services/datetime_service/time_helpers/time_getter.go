package time_helpers

//
//import time
import "time"

//from time import strftime
//from time import gmtime

//def now_time_as_string() \
//-> str:
func NowTimeAsString() string {

	utc := time.Now().UTC()

	//now_as_datetime_string = \
	//strftime(
	//"%Y-%m-%d %H:%M:%S",
	//gmtime())

	NowAsDataTimeString :=
		utc.Format("2006-01-02 15:04:05")

	//return \
	//now_as_datetime_string
	return NowAsDataTimeString

}

//def now_time_as_string_for_files() \
//-> str:
func NowTimeAsStringForFiles() string {

	//return \
	//strftime(
	//"%Y_%m_%d_%H_%M_%S",
	//gmtime())

	utc := time.Now().UTC()
	return utc.Format("2006_01_02_15_04_05")

}

//def time_as_string_yyyymmddhhmm(
//structured_time: time.struct_time) \
//-> str:
//formatted_string = \
//strftime(
//'%Y%m%d%H%M',
//structured_time)
//
//return \
//formatted_string

//def now_time_as_string_yyyymmddhhmm() \
//-> str:
//time_now = \
//gmtime()
//
//formatted_string = \
//time_as_string_yyyymmddhhmm(
//structured_time=time_now)
//
//return \
//formatted_string

//def __get_formatted_string_from_structured_time(
//format: str,
//structured_time: time.struct_time) -> str:
//formatted_string = \
//strftime(
//format,
//structured_time)
//
//return \
//formatted_string
