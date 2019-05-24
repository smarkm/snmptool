package cmd

import (
	"log"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//Receiver command
type Receiver struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *Receiver) Run(args []string) int {
	rs := 0
	ln := len(args)
	switch ln {
	case 0:
		log.Println("will start defualt trap receiver ...")
		r := &snmp.TrapReciver{Address: "127.0.0.1:162"}
		defer r.Close()
		r.Start()
	case 1:
		r := &snmp.TrapReciver{Address: args[0] + ":162"}
		defer r.Close()
		r.Start()
	case 2:
		port := args[1]
		r := &snmp.TrapReciver{Address: args[0] + ":" + port}
		defer r.Close()
		r.Start()
	default:
		c.UI.Output(c.Help())
		return 0
	}

	return rs
}

//Synopsis Synopsis information
func (c *Receiver) Synopsis() string {
	return "Start an trap receiver"
}

//Help Help information
func (c *Receiver) Help() string {
	return "Please use [ip] [port] to start an trap receiver"
}
