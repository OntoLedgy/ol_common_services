package operating_system_service

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ApplicationRunner struct {
	commandString               string
	commandArguments            string
	commandEnvironmentDrive     string
	commandEnvironmentDirectory string
}

func (applicationRunner *ApplicationRunner) RunCommand() {

	var stdBuffer bytes.Buffer

	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	args := strings.Split(" /C "+applicationRunner.commandArguments, " ")

	commandHandler := exec.Command(applicationRunner.commandString, args...)

	commandHandler.Dir = filepath.Join(applicationRunner.commandEnvironmentDrive, applicationRunner.commandEnvironmentDirectory)

	commandHandler.Stdout = mw

	commandHandler.Stderr = mw

	if err := commandHandler.Run(); err != nil {
		fmt.Printf("error:%s", err)
		log.Panic(err)
	}

	fmt.Printf("output: \n%s", stdBuffer.String())
}
