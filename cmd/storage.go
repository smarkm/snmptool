package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/smarkm/snmptool/snmp"
)

//Storage command
type Storage struct {
	//UI extend
	UI cli.Ui
}

//Run execute functioin
func (c *Storage) Run(args []string) (rs int) {
	ln := len(args)
	ip := ""
	communit := "public"
	switch ln {
	case 1:
		ip = args[0]
	case 2:
		communit = args[1]
	case 3:
	default:
		c.UI.Output(c.Help())
		return 0
	}
	items, err := snmp.GetWinStorage(ip, communit)
	if err != nil {
		log.Println(err)
	} else {
		c.UI.Output("Unit\t Size\t Used\t Desc")
		for _, item := range items {
			d := []string{strconv.Itoa(item.Units), strconv.FormatInt(item.Size, 10), strconv.FormatInt(item.Used, 10), item.Descr}
			c.UI.Output(strings.Join(d, "\t"))
		}
		c.UI.Output("Total: " + strconv.Itoa(len(items)) + " rows")
	}
	return
}

//Synopsis Synopsis information
func (c *Storage) Synopsis() string {
	return "Show basic storate information"
}

//Help Help information
func (c *Storage) Help() string {
	return c.Synopsis()
}
