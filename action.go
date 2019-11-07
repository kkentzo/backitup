package main

type Action struct {
	Command  string
	Artifact string
	Target   *Target
}
