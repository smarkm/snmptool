package snmp

import (
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//LLDP command
type LLDP struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *LLDP) Run(args []string) int {
	rs := 0
	ln := len(args)
	communit := "public"
	switch ln {
	case 1:
	case 2:
		communit = args[1]
	case 3:
	default:
		c.UI.Output(c.Help())
		return 0
	}
	items, err := snmp.GetLLdpLocalTable(args[0], communit)
	if err != nil {
		log.Println(err)
	} else {
		c.UI.Output("LldpLocPortNum\t LldpLocPortID\t LldpLocPortIDSubtype\t LldpLocPortDesc")
		for _, item := range items {
			d := []string{item.LldpLocPortNum, item.LldpLocPortID, item.LldpLocPortIDSubtype, item.LldpLocPortDesc}
			c.UI.Output(strings.Join(d, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(items)) + " rows")
	}
	return rs
}

//Synopsis Synopsis information
func (c *LLDP) Synopsis() string {
	return "get lldp"
}

//Help Help information
func (c *LLDP) Help() string {
	return "get lldp"
}

//LLDPRem command
type LLDPRem struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *LLDPRem) Run(args []string) int {
	rs := 0
	ln := len(args)
	communit := "public"
	switch ln {
	case 1:
	case 2:
		communit = args[1]
	case 3:
	default:
		c.UI.Output(c.Help())
		return 0
	}
	items, err := snmp.GetLLdpRemTable(args[0], communit)
	if err != nil {
		log.Println(err)
	} else {
		c.UI.Output("LLdpRemLocalPortNum\t LldpRLLdpRemSysNameemPortID\t LLdpRemPortID\t LLdpRemPortDesc")
		for _, item := range items {
			d := []string{item.LLdpRemLocalPortNum, item.LLdpRemSysName, item.LLdpRemPortID, item.LLdpRemPortDesc}
			c.UI.Output(strings.Join(d, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(items)) + " rows")
	}
	return rs
}

//Synopsis Synopsis information
func (c *LLDPRem) Synopsis() string {
	return "get lldp remote info"
}

//Help Help information
func (c *LLDPRem) Help() string {
	return "get lldp remote info"
}
