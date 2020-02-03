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

// iftableCmd represents the iftable command
var iftableCmd = &cobra.Command{
	Use:   "iftable",
	Short: "Show Iftable brief information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ports, err := snmp.GetPortsInformation(IP, Community)
		if err != nil {
			log.Println(err)
		} else {
			header := []string{"Index", "Admin", "Oper", "Name", "Speed"}
			fmt.Println(strings.Join(header, "\t"))
			for _, p := range ports {
				fmt.Println(strings.Join([]string{strconv.Itoa(p.Index), p.AdminStr(), p.OperStr(), p.Name, strconv.FormatInt(p.Speed, 10)}, "\t"))
			}
			fmt.Println("Total: " + strconv.Itoa(len(ports)) + " rows")
		}

	},
}

func init() {
	rootCmd.AddCommand(iftableCmd)
	UseGlobleFlags(iftableCmd)

}
