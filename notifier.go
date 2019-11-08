package main

type Notifier struct {
	Slack string
}

func (n *Notifier) SendReport(errors []error) error {
	return nil
}
