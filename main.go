package main

import (
	"fmt"
	"os"

	"github.com/smarkm/snmptool/cmd"

	"github.com/mitchellh/cli"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("snmp", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"sys": func() (cli.Command, error) {
			return &snmp.System{UI: ui}, nil
		},
		"interface": func() (cli.Command, error) {
			return &snmp.Interface{UI: ui}, nil
		},
		"lldp": func() (cli.Command, error) {
			return &snmp.LLDP{UI: ui}, nil
		},
		"lldprem": func() (cli.Command, error) {
			return &snmp.LLDPRem{UI: ui}, nil
		},
		"trapreceiver": func() (cli.Command, error) {
			return &snmp.Receiver{UI: ui}, nil
		},
	}

	_, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

}
