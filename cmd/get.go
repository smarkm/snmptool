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

	"github.com/smarkm/snmptool/snmp"
	"github.com/smarkm/snmptool/snmp/util"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Execute SNMP GET function",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		target := ParseOIDName(oid)
		rs, err := snmp.GetOne(IP, Community, target, ParseSNMPVer())
		if err != nil {
			util.HandleError(err)
			return
		}
		fmt.Println(snmp.GetSnmpString(rs))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&oid, "oid", "o", "", "target oid")
	getCmd.MarkFlagRequired("oid")

}
