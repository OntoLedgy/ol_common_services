package testing

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/datetime_service/time_helpers"
	"testing"
)

func TestDataTimeForFile(t *testing.T) {

	fmt.Println(time_helpers.NowTimeAsStringForFiles())

}

func TestDataTimeString(t *testing.T) {

	fmt.Println(time_helpers.NowTimeAsString())

}
