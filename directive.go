package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

type Directive struct {
	Command string
	Output  string
}

func (d *Directive) Execute() error {
	tokens := strings.Split(d.Command, " ")
	cmd := exec.Command(tokens[0], tokens[1:len(tokens)]...)
	// setup the output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	// run the command
	if err := cmd.Start(); err != nil {
		return err
	}
	if d.Output != "" {
		b, err := ioutil.ReadAll(stdout)
		if err != nil {
			return err
		}
		// grab the output and write it to the file
		if err := ioutil.WriteFile(d.Output, b, 0644); err != nil {
			return err
		}
	}
	return nil
}
