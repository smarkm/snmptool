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

	g "github.com/soniah/gosnmp"
	"github.com/spf13/cobra"
)

var (
	sourceIP  string
	eOID      string //enterprise OID
	genTraps  = make(map[string]string, 0)
	namedTrap string
)

// sendtrapCmd represents the sendtrap command
var sendtrapCmd = &cobra.Command{
	Use:   "sendtrap",
	Short: "Send trap",
	Long: `This a basic test tool currently only support coldStart,warmStart,linkDown,linkUp,aFailure(authenticationFailure),ep:
snmptool sendtrap -i target-ip -p target-port [-v 1|2] -t [coldStart|warmStart|linkDown|linkUp|aFailure]`,
	Run: func(cmd *cobra.Command, args []string) {
		g.Default.Target = IP
		g.Default.Port = port
		g.Default.Version = ParseSNMPVer()

		err := g.Default.Connect()
		if err != nil {
			fmt.Println(err.Error())
		}
		defer g.Default.Conn.Close()

		pdu := g.SnmpPDU{
			Name:  genTraps[namedTrap],
			Type:  g.OctetString,
			Value: namedTrap,
		}

		trap := g.SnmpTrap{
			Variables:    []g.SnmpPDU{pdu},
			Enterprise:   eOID,
			AgentAddress: sourceIP,
			GenericTrap:  0,
			SpecificTrap: 0,
			Timestamp:    300,
		}

		_, err = g.Default.SendTrap(trap)
		if err != nil {
			log.Fatalf("SendTrap() err: %v", err)
		}
	},
}

func init() {
	genTraps["coldStart"] = ".1.3.6.1.6.3.1.1.5.1"
	genTraps["warmStart"] = ".1.3.6.1.6.3.1.1.5.2"
	genTraps["linkDown"] = ".1.3.6.1.6.3.1.1.5.3"
	genTraps["linkUp"] = ".1.3.6.1.6.3.1.1.5.4"
	genTraps["aFailure"] = ".1.3.6.1.6.3.1.1.5.5"
	rootCmd.AddCommand(sendtrapCmd)
	sendtrapCmd.Flags().StringVarP(&namedTrap, "named trap", "t", "warmStart", "Named trap:coldStart|warmStart|linkDown|linkUp|aFailure")
	sendtrapCmd.Flags().StringVarP(&sourceIP, "source  ip", "s", "", "AgentAddress")

}
