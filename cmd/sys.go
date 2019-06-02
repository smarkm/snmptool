package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
	"github.com/smarkm/snmptool/snmp/util"
)

//System command
type System struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *System) Run(args []string) int {
	ip, community, err := ParseIPAndCommunity(args, 1)
	if err != nil {
		c.UI.Output(err.Error())
		c.UI.Output(c.Help())
		return 0
	}

	d, err := snmp.GetSystem(ip, community)
	if err != nil {
		log.Println(err)
		return 1
	} else {
		data := []string{"sysName: " + d.Name, "sysDescr: " + d.Desc, "sysObjectID: " + d.OId + " (" + util.GetDeviceType(d.OId) + ")", "sysContract: " + d.Contract,
			"sysLocation: " + d.Location, "sysServices: " + d.Services, "sysUpTime: " + d.UpTime}
		fmt.Println(strings.Join(data, "\n"))
	}
	return 0
}

//Synopsis Synopsis information
func (c *System) Synopsis() string {
	return "Show system information"
}

//Help Help information
func (c *System) Help() string {
	return "Usage: ... sys ip community [snmpversion]"
}
