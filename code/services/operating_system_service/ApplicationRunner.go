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
	CommandString               string
	CommandArguments            string
	CommandEnvironmentDrive     string
	CommandEnvironmentDirectory string
}

func (applicationRunner *ApplicationRunner) RunCommand() io.Writer {

	var stdBuffer bytes.Buffer

	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	args := strings.Split(
		" /C "+applicationRunner.CommandArguments,
		" ")

	commandHandler := exec.Command(
		applicationRunner.CommandString, args...)

	commandHandler.Dir = filepath.Join(
		applicationRunner.CommandEnvironmentDrive+":\\",
		applicationRunner.CommandEnvironmentDirectory)

	commandHandler.Stdout = mw

	commandHandler.Stderr = mw

	if err := commandHandler.Run(); err != nil {
		fmt.Printf("error:%s", err)
		log.Panic(err)
	}

	fmt.Printf("output: \n%s", stdBuffer.String())

	return mw
}
