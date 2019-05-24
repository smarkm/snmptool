package main

import (
	"fmt"
	"os"

	_ "github.com/go-bindata/go-bindata"
	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/cmd"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("snmp", "pre-0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"sys": func() (cli.Command, error) {
			return &cmd.System{UI: ui}, nil
		}, "storage": func() (cli.Command, error) {
			return &cmd.Storage{UI: ui}, nil
		}, "interface": func() (cli.Command, error) {
			return &cmd.Interface{UI: ui}, nil
		}, "ipaddr": func() (cli.Command, error) {
			return &cmd.IPAddress{UI: ui}, nil
		}, "lldp": func() (cli.Command, error) {
			return &cmd.LLDP{UI: ui}, nil
		}, "lldprem": func() (cli.Command, error) {
			return &cmd.LLDPRem{UI: ui}, nil
		}, "ospf": func() (cli.Command, error) {
			return &cmd.OSPF{UI: ui}, nil
		}, "ospfnbr": func() (cli.Command, error) {
			return &cmd.OSPFNbr{UI: ui}, nil
		}, "bgp": func() (cli.Command, error) {
			return &cmd.BGP{UI: ui}, nil
		}, "bgppeer": func() (cli.Command, error) {
			return &cmd.BGPPeer{UI: ui}, nil
		}, "trapreceiver": func() (cli.Command, error) {
			return &cmd.Receiver{UI: ui}, nil
		},
	}

	_, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

}
