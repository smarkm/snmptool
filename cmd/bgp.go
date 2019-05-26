package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

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
	items, err := snmp.GetBGPPeerTable(ip, communit)
	if err != nil {
		log.Println(err)
	} else {
		c.UI.Output("LocalAddr\t LocalPort\t RemoteAddr\t  RemotePort\t RemoteAS\t PeerState\t")
		for _, item := range items {
			d := []string{item.LocalAddr, strconv.Itoa(item.LocalPort), item.RemoteAddr, strconv.Itoa(item.RemotePort), strconv.Itoa(item.RemoteAS), item.PeerStateStr()}
			c.UI.Output(strings.Join(d, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(items)) + " rows")
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
