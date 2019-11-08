package main

import (
	"fmt"
	"os"
)

func Run(config *Config) error {
	errors := []error{}
	for _, action := range config.Actions {
		if err := action.Execute(); err != nil {
			errors = append(errors, err)
			continue
		}
		if err := action.Backup(); err != nil {
			errors = append(errors, err)
		}
	}

	if config.Notifier != nil {
		err := config.Notifier.SendReport(errors)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	config, err := NewConfigFromFile("demo.yml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := Run(config); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
