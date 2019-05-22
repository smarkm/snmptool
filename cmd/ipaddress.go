package cmd

import "github.com/mitchellh/cli"

//IPAddress command
type IPAddress struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *IPAddress) Run(args []string) int {
	return 0
}

//Synopsis Synopsis information
func (c *IPAddress) Synopsis() string {
	return "Show the ip address table"
}

//Help Help information
func (c *IPAddress) Help() string {
	return c.Synopsis()
}
