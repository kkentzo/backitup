package main

import (
	"fmt"
	"os"
)

func main() {
	c, err := NewConfigFromFile("demo.yml")
	if err != nil {
		os.Exit(1)
	}
	// execute the commands
	errors := []error{}
	for _, action := range c.Actions {
		if err := action.Execute(); err != nil {
			errors = append(errors, err)
			continue
		}
		if err := action.Backup(); err != nil {
			errors = append(errors, err)
		}
	}

	if c.Notifications != nil {
		err := c.Notifications.SendReport(errors)
		if err != nil {
			fmt.Printf("Notifiaction error: %s", err.Error())
		}
	}
}
