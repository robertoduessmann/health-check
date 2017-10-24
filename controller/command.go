package controller

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func ExecCommand(command string) (out bytes.Buffer) {
	commandWithArgs := strings.Split(command, " ")
	cmd := exec.Command(commandWithArgs[0], commandWithArgs[1:]...)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}
