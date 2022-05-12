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

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Show storage biref information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := snmp.GetWinStorage(snmp.NewSNMP(IP, Community, ParseSNMPVer()))
		if err != nil {
			util.HandleError(err)
		} else {
			fmt.Println("Unit\t Size\t Used\t Desc")
			for _, item := range items {
				d := []string{strconv.Itoa(item.Units), strconv.FormatInt(item.Size, 10), strconv.FormatInt(item.Used, 10), item.Descr}
				fmt.Println(strings.Join(d, "\t"))
			}
			fmt.Println("Total: " + strconv.Itoa(len(items)) + " rows")
		}
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)

}
