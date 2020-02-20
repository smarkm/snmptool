package snmp

import (
	"log"
	"strconv"

	g "github.com/soniah/gosnmp"
)

//Oid of ipAddTable[.1.3.6.1.2.1.4.20]
const (
	OIDIpAdEntAddr    = ".1.3.6.1.2.1.4.20.1.1"
	OIDIpAdEntIfIndex = ".1.3.6.1.2.1.4.20.1.2"
	OIDIpAdEntNetMask = ".1.3.6.1.2.1.4.20.1.3"
)

//IPAddr Entity
type IPAddr struct {
	IP      string
	IfIndex int
	Netmask string
	IfDesc  string
}

//GetIPAddrTable get loclTable
func GetIPAddrTable(ip string, communit string, ver g.SnmpVersion) (ips []*IPAddr, err error) {
	oids := []string{OIDIpAdEntAddr, OIDIpAdEntIfIndex, OIDIpAdEntNetMask}
	tableRows, err := GetTable(ip, communit, ver, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		item := new(IPAddr)
		item.IP = GetSnmpString(row[OIDIpAdEntAddr])
		item.IfIndex = GetSnmpInt(row[OIDIpAdEntIfIndex])
		item.Netmask = GetSnmpString(row[OIDIpAdEntNetMask])
		ips = append(ips, item)
	}
	for _, ipa := range ips {
		v, err := GetOne(ip, communit, IfDescr+"."+strconv.Itoa(ipa.IfIndex), ver)
		if err != nil {
			log.Println(err.Error())
		} else {
			ipa.IfDesc = GetSnmpString(v)
		}
	}
	return
}
