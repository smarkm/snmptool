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
	"github.com/soniah/gosnmp"
	"github.com/spf13/cobra"
)

// walkCmd represents the walk command
var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: "Execute SNMP WALK",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := snmp.NewSNMP(IP, Community)
		err := s.Connect()
		if err != nil {
			return
		}
		defer s.Conn.Close()
		target := ParseOIDName(oid)
		err = s.Walk(target, func(data gosnmp.SnmpPDU) error {
			fmt.Println(snmp.GetSnmpString(data))
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(walkCmd)
	UseGlobleFlags(walkCmd)
	walkCmd.Flags().StringVarP(&oid, "oid", "o", "", "root oid")
	getCmd.MarkFlagRequired("oid")
}
