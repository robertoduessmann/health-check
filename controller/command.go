package controller

import (
	"bytes"
	"log"
	"os/exec"
)

func ExecCommand(command string) (out bytes.Buffer) {
	cmd := exec.Command(command)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}
