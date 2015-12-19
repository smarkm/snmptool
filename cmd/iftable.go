package snmp

//this file will be deprected and remove [Remove]
import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//IfTable command
type IfTable struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *IfTable) Run(args []string) (rs int) {
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
	ports, err := snmp.GetPortsInformation(args[0], communit)
	if err != nil {
		log.Println(err)
		rs = 1
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
func (c *IfTable) Synopsis() string {
	return "get ifTable"
}

//Help Help information
func (c *IfTable) Help() string {
	return "get ifTable"
}
