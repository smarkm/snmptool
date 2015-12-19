package snmp

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//Interface show interface brief base index
type Interface struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *Interface) Run(args []string) (rs int) {
	ln := len(args)
	communit := "public"

	if ln < 3 {
		c.UI.Output(c.Help())
		return
	}

	ip, communit, index := args[0], args[1], args[2]

	if "all" == index {
		c.showAllPortsInformation(ip, communit)
		return
	}

	c.showOnePortInformation(ip, communit, index)

	return
}
func (c *Interface) showOnePortInformation(ip, community, index string) (err error) {
	port, err := snmp.GetPortInformation(ip, community, index)
	if err != nil {
		log.Println(err)
	} else {
		header := []string{"Index", "Admin", "Oper", "Name", "Speed(bits)"}
		c.UI.Output(strings.Join(header, "\t"))
		fmt.Println(strings.Join([]string{strconv.Itoa(port.Index), port.AdminStr(), port.OperStr(), port.Name, strconv.FormatInt(port.Speed, 10)}, "\t"))
	}
	return
}

func (c *Interface) showAllPortsInformation(ip, community string) (err error) {
	ports, err := snmp.GetPortsInformation(ip, community)
	if err != nil {
		log.Println(err)
		return
	} else {
		header := []string{"Index", "Admin", "Oper", "Name", "Speed"}
		c.UI.Output(strings.Join(header, "\t"))
		for _, p := range ports {
			fmt.Println(strings.Join([]string{strconv.Itoa(p.Index), p.AdminStr(), p.OperStr(), p.Name, strconv.FormatInt(p.Speed, 10)}, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(ports)) + " rows")
	}
	return
}

//Synopsis Synopsis information
func (c *Interface) Synopsis() string {
	return "show interface brief base index"
}

//Help Help information
func (c *Interface) Help() string {
	return "args should be : ip communit index"
}
