package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//CDP command
type CDP struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *CDP) Run(args []string) (rs int) {
	ip, community, err := ParseIPAndCommunity(args, 1)
	if err != nil {
		c.UI.Output(err.Error())
		c.UI.Output(c.Help())
		return 0
	}
	items, err := snmp.GetCDPTable(ip, community)
	if err != nil {
		log.Println(err)
	} else {
		c.UI.Output("IfIndex\t Enable\t MsgInterval\t Group\t Port\t Name")
		for _, item := range items {
			d := []string{i2S(item.IfIndex), item.ParseEnable(), i2S(item.MsgInterval), i2S(item.Group), i2S(item.Port), item.Name}
			c.UI.Output(strings.Join(d, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(items)) + " rows")
	}

	return
}

//Synopsis Synopsis information
func (c *CDP) Synopsis() string {
	return "Show cisco CDP information"
}

//Help Help information
func (c *CDP) Help() string {
	return c.Synopsis()
}
