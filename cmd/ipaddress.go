package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//IPAddress command
type IPAddress struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *IPAddress) Run(args []string) (rs int) {
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
	items, err := snmp.GetIpAddrTable(ip, communit)
	if err != nil {
		log.Println(err)
	} else {
		c.UI.Output("IP\t Netmask\t Index\t Descr")
		for _, item := range items {
			d := []string{item.IP, item.Netmask, strconv.Itoa(item.IfIndex), item.IfDesc}
			c.UI.Output(strings.Join(d, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(items)) + " rows")
	}
	return
}

//Synopsis Synopsis information
func (c *IPAddress) Synopsis() string {
	return "Show the ip address table"
}

//Help Help information
func (c *IPAddress) Help() string {
	return c.Synopsis()
}
