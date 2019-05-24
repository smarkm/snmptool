package cmd

import "github.com/mitchellh/cli"

//OSPF command
type OSPF struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *OSPF) Run(args []string) (rs int) {
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
func (c *OSPF) Synopsis() string {
	return "Show ospf interface information"
}

//Help Help information
func (c *OSPF) Help() string {
	return c.Synopsis()
}

//OSPFNbr command
type OSPFNbr struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *OSPFNbr) Run(args []string) (rs int) {
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
func (c *OSPFNbr) Synopsis() string {
	return "Show ospf peer information"
}

//Help Help information
func (c *OSPFNbr) Help() string {
	return c.Synopsis()
}
