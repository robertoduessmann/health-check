package controller

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func HealtCheck(w http.ResponseWriter, r *http.Request) {
	respose := execCommand("free")
	fmt.Fprintf(w, respose.String())
}

func execCommand(command string) (out bytes.Buffer) {
	cmd := exec.Command(command)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}
