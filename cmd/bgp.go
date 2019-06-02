package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//BGPPeer command
type BGPPeer struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *BGPPeer) Run(args []string) (rs int) {
	ip, community, err := ParseIPAndCommunity(args, 1)
	if err != nil {
		c.UI.Output(err.Error())
		c.UI.Output(c.Help())
		return 0
	}
	items, err := snmp.GetBGPPeerTable(ip, community)
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
