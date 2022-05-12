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
	tableRows, _ := GetTable(local, communit, snmpver, oids)
	log.Println(len(tableRows))
	fmt.Println(len(tableRows))
}
func TestGetPortInformation(t *testing.T) {

}

func TestGetIPTable(t *testing.T) {
	ips, err := GetIPTable(local, communit, snmpver)
	for index, ip := range ips {
		log.Println(index, ip, err)
	}
}

func TestWinCPU(t *testing.T) {
	table, _ := GetWinStorage(local, communit, snmpver)
	for _, item := range table {
		fmt.Println(item)
	}
}

func TestHrProcessorLoad(t *testing.T) {
	table, _ := GetHrProcessorLoad(local, communit, snmpver)
	for _, item := range table {
		fmt.Println(item)
	}
}

func TestGetIPAddrTable(t *testing.T) {
	table, _ := GetIPAddrTable(local, communit, snmpver)
	for _, item := range table {
		fmt.Println(item)
		fmt.Println("")
	}
}

func TestGetHostName(t *testing.T) {
	log.Println(GetHostName("172.16.2.2", "public", snmpver))
}

func TestGetIPForaordTable(t *testing.T) {
	items, err := GetIPForwardTable(local, communit, snmpver)
	for _, item := range items {
		log.Println(item)
	}
	util.HandleError(err)
}
