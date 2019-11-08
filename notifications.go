package main

type Notifications struct {
	Slack string
}

func (n *Notifications) SendReport(errors []error) error {
	return nil
}
