package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Command(t *testing.T) {
	fname := "list.txt"
	d := Directive{Command: "ls -l", Output: fname}
	defer os.Remove(fname)

	assert.Nil(t, d.Execute())

	file, err := os.Open(fname)
	defer file.Close()
	assert.Nil(t, err)

	b, err := ioutil.ReadAll(file)
	assert.Nil(t, err)
	contents := string(b)
	assert.Contains(t, contents, "total")
	assert.Contains(t, contents, "directive.go")
	assert.Contains(t, contents, "directive_test.go")
}
