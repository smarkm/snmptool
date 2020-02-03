/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/smarkm/snmptool/snmp"
	"github.com/spf13/cobra"
)

var index string

// interfaceCmd represents the interface command
var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "Show interface biref information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(index)
		if "all" == index {
			showAllPortsInformation(IP, Community)
		}

		showOnePortInformation(IP, Community, index)

	},
}

func init() {
	rootCmd.AddCommand(interfaceCmd)
	UseGlobleFlags(interfaceCmd)
	interfaceCmd.Flags().StringVarP(&index, "index", "", "", "ifindex")
	interfaceCmd.MarkFlagRequired("index")
}
func showOnePortInformation(ip, community, index string) (err error) {
	port, err := snmp.GetPortInformation(ip, community, index)
	if err != nil {
		log.Println(err)
	} else {
		header := []string{"Index", "Admin", "Oper", "Name", "Speed(bits)"}
		fmt.Println(strings.Join(header, "\t"))
		fmt.Println(strings.Join([]string{strconv.Itoa(port.Index), port.AdminStr(), port.OperStr(), port.Name, strconv.FormatInt(port.Speed, 10)}, "\t"))
	}
	return
}

func showAllPortsInformation(ip, community string) (err error) {
	ports, err := snmp.GetPortsInformation(ip, community)
	if err != nil {
		log.Println(err)
		return
	} else {
		header := []string{"Index", "Admin", "Oper", "Name", "Speed"}
		fmt.Println(strings.Join(header, "\t"))
		for _, p := range ports {
			fmt.Println(strings.Join([]string{strconv.Itoa(p.Index), p.AdminStr(), p.OperStr(), p.Name, strconv.FormatInt(p.Speed, 10)}, "\t"))
		}
		fmt.Println("Total: " + strconv.Itoa(len(ports)) + " rows")
	}
	return
}
