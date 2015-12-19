package snmp

import (
	"github.com/mitchellh/cli"
)

//SNMP command
type SNMP struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *SNMP) Run(args []string) (rs int) {
	snmpCli := cli.NewCLI("snmp", "")
	snmpCli.Args = args

	snmpCli.Commands = map[string]cli.CommandFactory{
		"sys": func() (cli.Command, error) {
			return &System{UI: c.UI}, nil
		},
		"interface": func() (cli.Command, error) {
			return &Interface{UI: c.UI}, nil
		},
		"lldp": func() (cli.Command, error) {
			return &LLDP{UI: c.UI}, nil
		},
		"lldprem": func() (cli.Command, error) {
			return &LLDPRem{UI: c.UI}, nil
		},
		"trapreceiver": func() (cli.Command, error) {
			return &Receiver{UI: c.UI}, nil
		},
	}

	if exitStatus, err := snmpCli.Run(); err != nil {
		c.UI.Error(err.Error())
		rs = exitStatus
	} else {
		rs = exitStatus
	}
	return
}

//Synopsis Synopsis information
func (c *SNMP) Synopsis() string {
	return "snmp tool"
}

//Help Help information
func (c *SNMP) Help() string {
	return "simple snmp tool"
}
