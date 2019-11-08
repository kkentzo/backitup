package main

import (
	"fmt"
)

type Action struct {
	Directive *Directive
	Artifact  *Artifact
	Target    *Target
}

// execute the action's command
func (a *Action) Execute() error {
	// skip alltogether?
	if a.Directive == nil {
		return nil
	}
	return a.Directive.Execute()
}

// copy the artifact to the target
func (a *Action) Backup() error {
	// skip alltogether?
	if a.Artifact == nil {
		return nil
	}
	// does the artifact exist?
	if !a.Artifact.Exists() {
		return fmt.Errorf("Artifact %s does not exist", a.Artifact.Path)
	}
	if err := a.Artifact.Copy(a.Target); err != nil {
		return err
	}
	return nil
}
