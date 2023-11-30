package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

const CLIPluginServerAddr = ":8020"

func executeCmd(prefix string) []byte {
	cmd := exec.Command("./plugin.cli")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	io.WriteString(stdin, prefix+"\n")

	output, err := io.ReadAll(stdout)
	if err != nil {
		panic(err)
	}

	return output
}

func ServeCLIPlugin() {
	http.HandleFunc("/cli", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, executeCmd(Name))
	})
	http.ListenAndServe(CLIPluginServerAddr, nil)
}
