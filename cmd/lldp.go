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

// lldpCmd represents the lldp command
var lldpCmd = &cobra.Command{
	Use:   "lldp",
	Short: "Show LLDP brief information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := snmp.GetLLdpLocalTable(IP, Community)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("LldpLocPortNum\t LldpLocPortID\t LldpLocPortIDSubtype\t LldpLocPortDesc")
			for _, item := range items {
				d := []string{item.LldpLocPortNum, item.LldpLocPortID, item.LldpLocPortIDSubtype, item.LldpLocPortDesc}
				fmt.Println(strings.Join(d, "\t"))
			}
			fmt.Println("Total: " + strconv.Itoa(len(items)) + " rows")
		}
	},
}

func init() {
	rootCmd.AddCommand(lldpCmd)
	UseGlobleFlags(lldpCmd)
}
