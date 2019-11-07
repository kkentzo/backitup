package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewFromFile(t *testing.T) {
	c, err := NewConfigFromFile("demo.yml")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.Actions))
	a := c.Actions[0]
	assert.Equal(t, "ls -l > list.txt", a.Command)
	assert.Equal(t, "list.txt", a.Artifact)
	assert.Equal(t, "s3://my-backup-bucket", a.Target.Bucket)
	assert.Equal(t, "lists", a.Target.Prefix)
}
