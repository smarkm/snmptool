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

// cdpCmd represents the cdp command
var cdpCmd = &cobra.Command{
	Use:   "cdp",
	Short: "Show CDP brief infromation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := snmp.GetCDPTable(snmp.NewSNMP(IP, Community, ParseSNMPVer()))
		if err != nil {
			util.HandleError(err)
		} else {
			fmt.Println("IfIndex\t Enable\t MsgInterval\t Group\t Port\t Name")
			for _, item := range items {
				d := []string{i2S(item.IfIndex), item.ParseEnable(), i2S(item.MsgInterval), i2S(item.Group), i2S(item.Port), item.Name}
				fmt.Println(strings.Join(d, "\t"))
			}
			fmt.Println("Total: " + strconv.Itoa(len(items)) + " rows")
		}
	},
}

func init() {
	rootCmd.AddCommand(cdpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cdpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cdpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
