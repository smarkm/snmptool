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
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/spf13/cobra"
)

var addParams bool
var delParams int
var defaultParamIndex int

// getCmd represents the get command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config default snmp params or management snmp params ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		var params []SNMPParams

		fileName := path.Join(dirname, ".snmptool")
		configFile, err := os.Open(fileName)
		if err != nil {
			public := SNMPParams{SNMPVersion: snmpver, Community: "public", Default: true}
			params = append(params, public)
			data, _ := json.Marshal(params)
			ioutil.WriteFile(fileName, data, 0644)
		}
		data, err := ioutil.ReadAll(configFile)
		err = json.Unmarshal(data, &params)
		if addParams {
			param := SNMPParams{SNMPVersion: snmpver, Community: Community,
				Username: UserName, Level: SecurityLevel,
				AuthProtocol: AuthProtocolStr, PrivProtocol: PrivProtocolStr, PrivPass: PrivPass}
			params = append(params, param)
			data, _ := json.Marshal(params)
			ioutil.WriteFile(fileName, data, 0644)
			log.Println("total ", len(params), " snmp params")
		}
		if delParams > -1 {
			nowSize := len(params)
			if nowSize <= delParams {
				log.Println("Invalid item index, available index 0~" + strconv.Itoa(nowSize-1))
				return
			}
			params[delParams] = params[nowSize-1]
			params = params[:nowSize-1]
			data, _ := json.Marshal(params)
			ioutil.WriteFile(fileName, data, 0644)
			log.Println("success remove one params,remain " + strconv.Itoa(len(params)) + " snmp params")
		}

		if defaultParamIndex > -1 {
			nowSize := len(params)
			if nowSize <= defaultParamIndex {
				log.Println("Invalid item index, available index 0~" + strconv.Itoa(nowSize-1))
				return
			}
			for i := 0; i < nowSize; i++ {
				if i == defaultParamIndex {
					params[i].Default = true
				} else {
					params[i].Default = false
				}
			}
			data, _ := json.Marshal(params)
			ioutil.WriteFile(fileName, data, 0644)
			log.Println("success set default params")
		}

		if !(addParams || delParams > -1 || defaultParamIndex > -1) {
			for i := 0; i < len(params); i++ {
				p := params[i]
				isDefaul := ""
				if p.Default {
					isDefaul = ", default"
				}
				switch p.SNMPVersion {
				case "2c":
					log.Println("Index: ", i, ",Version: ", p.SNMPVersion, ",Community: "+p.Community, isDefaul)
				case "3":
					log.Println(p)
				}

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&addParams, "save", "s", false, "Add or update snmp params in to local config for ref")
	configCmd.Flags().IntVarP(&delParams, "remove", "r", -1, "Remove snmp params in local config file")
	configCmd.Flags().IntVarP(&defaultParamIndex, "default", "d", -1, "Set default snmp params")

}

func DefaulSNMPParams() SNMPParams {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	var params []SNMPParams

	fileName := path.Join(dirname, ".snmptool")
	configFile, err := os.Open(fileName)
	if err != nil {
		public := SNMPParams{SNMPVersion: snmpver, Community: "public", Default: true}
		params = append(params, public)
		data, _ := json.Marshal(params)
		ioutil.WriteFile(fileName, data, 0644)
	}
	data, err := ioutil.ReadAll(configFile)
	err = json.Unmarshal(data, &params)
	var p SNMPParams
	for _, v := range params {
		if v.Default {
			p = v
		}
	}
	return p
}

type SNMPParams struct {
	SNMPVersion  string `json:"version"`
	Community    string `json:"community"`
	Username     string `json:"username"`
	Level        string `json':"level"`
	AuthProtocol string `json:"authProtocol"`
	AuthPass     string `json:"authPass"`
	PrivProtocol string `json:"privProtocol"`
	PrivPass     string `json:"privPass"`
	Default      bool   `json:"default"`
}
