package snmp

import (
	"fmt"
	"log"
	"testing"
)

var local = "127.0.0.1"
var communit = "public"

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
	tableRows, _ := GetTable(local, communit, oids)
	fmt.Println(len(tableRows))
}
func TestGetPortInformation(t *testing.T) {

}

func TestGetIPTable(t *testing.T) {
	ips, err := GetIPTable(local, communit)
	for index, ip := range ips {
		log.Println(index, ip, err)
	}
}

func TestWinCPU(t *testing.T) {
	table, _ := GetWinStorage(local, communit)
	for _, item := range table {
		fmt.Println(item)
	}
}

func TestHrProcessorLoad(t *testing.T) {
	table, _ := GetHrProcessorLoad(local, communit)
	for _, item := range table {
		fmt.Println(item)
	}
}
