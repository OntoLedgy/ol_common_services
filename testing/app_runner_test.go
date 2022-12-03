package testing

import (
	"github.com/OntoLedgy/ol_common_services/code/services/operating_system_service"
	"testing"
)

func TestAppRunner(t *testing.T) {

	var applicationrunner = operating_system_service.ApplicationRunner{
		"cmd",
		"java org.antlr.v4.Tool",
		"D",
		"Apps"}

	applicationrunner.RunCommand()

}
