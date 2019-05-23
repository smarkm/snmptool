package snmp

import (
	"log"
	"strconv"
)

//Oid of ipAddTable[.1.3.6.1.2.1.4.20]
const (
	OIDIpAdEntAddr    = ".1.3.6.1.2.1.4.20.1.1"
	OIDIpAdEntIfIndex = ".1.3.6.1.2.1.4.20.1.2"
	OIDIpAdEntNetMask = ".1.3.6.1.2.1.4.20.1.3"
)

//IpAddr Entity
type IpAddr struct {
	IP      string
	IfIndex int
	Netmask string
	IfDesc  string
}

//GetLLdpLocalTable get loclTable
func GetIpAddrTable(ip string, communit string) (ips []*IpAddr, err error) {
	oids := []string{OIDIpAdEntAddr, OIDIpAdEntIfIndex, OIDIpAdEntNetMask}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		item := new(IpAddr)
		item.IP = GetSnmpString(row[OIDIpAdEntAddr])
		item.IfIndex = GetSnmpInt(row[OIDIpAdEntIfIndex])
		item.Netmask = GetSnmpString(row[OIDIpAdEntNetMask])
		ips = append(ips, item)
	}
	for _, ipa := range ips {
		v, err := GetOne(ip, communit, IfDescr+"."+strconv.Itoa(ipa.IfIndex))
		if err != nil {
			log.Println(err.Error())
		} else {
			ipa.IfDesc = GetSnmpString(v)
		}
	}
	return
}
