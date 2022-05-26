package testing

import (
	"fmt"
	"github.com/gen2brain/dlgs"

	"testing"
)

func TestEntryBox(t *testing.T) {

	entry_reading, entry_reading2, entry_reading3 := dlgs.Entry("test", "tesx", "hhh")

	fmt.Printf("%s \n %v \n %s, ", entry_reading, entry_reading2, entry_reading3)

}

func TestEntryBo1(t *testing.T) {

	entry_reading, entry_reading2, entry_reading3 := dlgs.Entry("test", "tesx", "hhh")

	fmt.Printf("%s \n %v \n %s, ", entry_reading, entry_reading2, entry_reading3)

	dlgs.Question("the question", "is it ok?", true)

}

func TestFile(t *testing.T) {

	dlgs.File("the file", "*.txt", true)
}
