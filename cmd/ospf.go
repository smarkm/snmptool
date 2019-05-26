package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//OSPF command
type OSPF struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *OSPF) Run(args []string) (rs int) {
	ln := len(args)
	ip := ""
	communit := "public"
	switch ln {
	case 1:
		ip = args[0]
	case 2:
		ip = args[0]
		communit = args[1]
	case 3:
	default:
		c.UI.Output(c.Help())
		return 0
	}
	items, err := snmp.GetOspfIfTable(ip, communit)
	if err != nil {
		log.Println(err)
	} else {
		c.UI.Output("IP Address\t AdminState\t State\t ")
		for _, item := range items {
			d := []string{item.OspfIfIPAddress, item.AdminStateStr(), item.StateStr()}
			c.UI.Output(strings.Join(d, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(items)) + " rows")
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
	UI cli.Ui
}

//Run execute functioin
func (c *OSPFNbr) Run(args []string) (rs int) {
	ln := len(args)
	ip := ""
	communit := "public"
	switch ln {
	case 1:
		ip = args[0]
	case 2:
		ip = args[0]
		communit = args[1]
	case 3:
	default:
		c.UI.Output(c.Help())
		return 0
	}
	items, err := snmp.GetOspfNbrTable(ip, communit)
	if err != nil {
		log.Println(err)
	} else {
		c.UI.Output("Nbr IP\t Nbr Router\t Nbr State\t ")
		for _, item := range items {
			d := []string{item.NbrIPAddress, item.NbrRtrID, item.NbrStateStr()}
			c.UI.Output(strings.Join(d, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(items)) + " rows")
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
