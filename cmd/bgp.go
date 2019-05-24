package cmd

import "github.com/mitchellh/cli"

//BGP command
type BGP struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *BGP) Run(args []string) (rs int) {
	ln := len(args)
	switch ln {
	case 1:

	case 2:
	default:
		c.UI.Output(c.Help())
		return rs
	}
	return
}

//Synopsis Synopsis information
func (c *BGP) Synopsis() string {
	return "Show BGP information"
}

//Help Help information
func (c *BGP) Help() string {
	return c.Synopsis()
}

//BGPPeer command
type BGPPeer struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *BGPPeer) Run(args []string) (rs int) {
	ln := len(args)
	switch ln {
	case 1:

	case 2:
	default:
		c.UI.Output(c.Help())
		return rs
	}
	return
}

//Synopsis Synopsis information
func (c *BGPPeer) Synopsis() string {
	return "Show BGP peer information"
}

//Help Help information
func (c *BGPPeer) Help() string {
	return c.Synopsis()
}
