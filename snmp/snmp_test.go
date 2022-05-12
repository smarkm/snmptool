package snmp

import (
	"fmt"
	"log"
	"testing"

	"github.com/smarkm/snmptool/snmp/util"

	"github.com/soniah/gosnmp"
)

var local = "127.0.0.1"
var communit = "public"
var snmpver = gosnmp.Version2c
var s = NewSNMP(local, communit, snmpver)

func TestSnmpClient_Get(t *testing.T) {
	// for {
	// 	time.Sleep(time.Second)
	// 	device, err := GetDeviceInfo("127.0.0.1", "public")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(device)
	// 	}
	// }
}

func TestSnmpGetTable(t *testing.T) {
	oids := []string{IfIndex, IfMtu}
	tableRows, _ := GetTable(s, oids)
	log.Println(len(tableRows))
	fmt.Println(len(tableRows))
}
func TestGetPortInformation(t *testing.T) {

}

func TestGetIPTable(t *testing.T) {
	ips, err := GetIPTable(s)
	for index, ip := range ips {
		log.Println(index, ip, err)
	}
}

func TestWinCPU(t *testing.T) {
	table, _ := GetWinStorage(s)
	for _, item := range table {
		fmt.Println(item)
	}
}

func TestHrProcessorLoad(t *testing.T) {
	table, _ := GetHrProcessorLoad(s)
	for _, item := range table {
		fmt.Println(item)
	}
}

func TestGetIPAddrTable(t *testing.T) {
	table, _ := GetIPAddrTable(s)
	for _, item := range table {
		fmt.Println(item)
		fmt.Println("")
	}
}

func TestGetHostName(t *testing.T) {
	log.Println(GetHostName(s))
}

func TestGetIPForaordTable(t *testing.T) {
	items, err := GetIPForwardTable(s)
	for _, item := range items {
		log.Println(item)
	}
	util.HandleError(err)
}
