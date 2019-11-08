package main

import "os"

type Artifact struct {
	Path string
	Keep bool
}

func (a *Artifact) Exists() bool {
	if _, err := os.Stat(a.Path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (a *Artifact) Copy(target *Target) error {
	// TODO: copy artifact to the target bucket
	return nil
}

func (a *Artifact) CleanUp() error {
	if a.Keep {
		return nil
	}
	// TODO: Remove the artifact from the filesystem
	return nil
}
