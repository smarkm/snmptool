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
	"strconv"
	"strings"

	"github.com/smarkm/snmptool/snmp"
	"github.com/smarkm/snmptool/snmp/util"
	"github.com/spf13/cobra"
)

// ipaddressCmd represents the ipaddress command
var ipaddressCmd = &cobra.Command{
	Use:   "ipaddress",
	Short: "Show ip address table",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := snmp.GetIPAddrTable(snmp.NewSNMP(IP, Community, ParseSNMPVer()))
		if err != nil {
			util.HandleError(err)
		} else {
			fmt.Println("IP\t Netmask\t Index\t Descr")
			for _, item := range items {
				d := []string{item.IP, item.Netmask, strconv.Itoa(item.IfIndex), item.IfDesc}
				fmt.Println(strings.Join(d, "\t"))
			}
			fmt.Println("Total: " + strconv.Itoa(len(items)) + " rows")
		}
	},
}

func init() {
	rootCmd.AddCommand(ipaddressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipaddressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipaddressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
