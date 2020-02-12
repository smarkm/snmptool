package snmp

import (
	"strconv"

	uc "github.com/smarkm/golang-underscore"
	"github.com/smarkm/snmptool/snmp/model"
)

//IfTable
const (
	IfIndex           = ".1.3.6.1.2.1.2.2.1.1"
	IfDescr           = ".1.3.6.1.2.1.2.2.1.2"
	IfType            = ".1.3.6.1.2.1.2.2.1.3"
	IfMtu             = ".1.3.6.1.2.1.2.2.1.4"
	IfSpeed           = ".1.3.6.1.2.1.2.2.1.5"
	IfPhysAddress     = ".1.3.6.1.2.1.2.2.1.6"
	IfAdminStatus     = ".1.3.6.1.2.1.2.2.1.7"
	IfOperStatus      = ".1.3.6.1.2.1.2.2.1.8"
	IfLastChange      = ".1.3.6.1.2.1.2.2.1.9"
	IfInOctets        = ".1.3.6.1.2.1.2.2.1.10"
	IfInUcastPkts     = ".1.3.6.1.2.1.2.2.1.11"
	IfInNUcastPkts    = ".1.3.6.1.2.1.2.2.1.12"
	IfInDiscards      = ".1.3.6.1.2.1.2.2.1.13"
	IfInErrors        = ".1.3.6.1.2.1.2.2.1.14"
	IfInUnknownProtos = ".1.3.6.1.2.1.2.2.1.15"
	IfOutOctets       = ".1.3.6.1.2.1.2.2.1.16"
	IfOutUcastPkts    = ".1.3.6.1.2.1.2.2.1.17"
	IfOutNUcastPkts   = ".1.3.6.1.2.1.2.2.1.18"
	IfOutDiscards     = ".1.3.6.1.2.1.2.2.1.19"
	IfOutErrors       = ".1.3.6.1.2.1.2.2.1.20"
	IfOutQLen         = ".1.3.6.1.2.1.2.2.1.21"
	IfSpecific        = ".1.3.6.1.2.1.2.2.1.22"
)

//ipAddrTable
const (
	ipAdEntAddr         = ".1.3.6.1.2.1.4.20.1.1"
	ipAdEntIfIndex      = ".1.3.6.1.2.1.4.20.1.2"
	ipAdEntNetMask      = ".1.3.6.1.2.1.4.20.1.3"
	ipAdEntBcastAddr    = ".1.3.6.1.2.1.4.20.1.4"
	ipAdEntReasmMaxSize = ".1.3.6.1.2.1.4.20.1.5"
)

func init() {
	OIDs["ifIndex"] = IfIndex
	OIDs["ifDescr"] = IfDescr
	OIDs["ifType"] = IfType
	OIDs["ifMtu"] = IfMtu
	OIDs["ifSpeed"] = IfSpeed
	OIDs["ifPhysAddress"] = IfPhysAddress
	OIDs["ifAdminStatus"] = IfAdminStatus
	OIDs["ifOperStatus"] = IfOperStatus
	OIDs["ifLastChange"] = IfLastChange
	OIDs["ifInOctects"] = IfInOctets
	OIDs["ifInUcastPkts"] = IfInUcastPkts
	OIDs["ifInNUcastPkts"] = IfInNUcastPkts
	OIDs["ifInDiscards"] = IfInDiscards
	OIDs["ifInErrors"] = IfInErrors
	OIDs["ifInUnknownProtos"] = IfInUnknownProtos
	OIDs["ifOutOctets"] = IfOutOctets
	OIDs["ifOutUcastPkts"] = IfOutUcastPkts
	OIDs["ifOutNUcastPkts"] = IfOutNUcastPkts
	OIDs["ifOutDiscards"] = IfOutDiscards
	OIDs["ifOutErrors"] = IfOutErrors
	OIDs["ifOutQLen"] = IfOutQLen
	OIDs["ifSpecific"] = IfSpeed
}

//GetPortsInformation get basic port information
func GetPortsInformation(ip string, communit string) (ports map[int]*model.Port, err error) {
	oids := []string{IfIndex, IfDescr, IfAdminStatus, IfOperStatus, IfMtu, IfSpeed, IfType, IfPhysAddress}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return ports, err
	}
	ports = make(map[int]*model.Port)
	for _, row := range tableRows {
		port := new(model.Port)
		port.Index = GetSnmpInt(row[IfIndex])
		port.Name = GetSnmpString(row[IfDescr])
		port.Descr = GetSnmpString(row[IfDescr])
		port.AdminStatus = GetSnmpInt(row[IfAdminStatus])
		port.OperStatus = GetSnmpInt(row[IfOperStatus])
		port.Mtu = GetSnmpInt(row[IfMtu])
		port.Speed = GetSnmpInt64(row[IfSpeed])
		port.Mac = GetSnmpMacString(row[IfPhysAddress])
		ports[port.Index] = port
	}
	return
}

//GetIPTable get basic port information
func GetIPTable(ip string, communit string) (ips map[int]string, err error) {
	oids := []string{ipAdEntAddr, ipAdEntIfIndex}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return ips, err
	}
	ips = make(map[int]string)
	for _, row := range tableRows {
		index := GetSnmpInt(row[ipAdEntIfIndex])
		ips[index] = GetSnmpString(row[ipAdEntAddr])
	}
	return
}

//GetPortsStatus get basic port information
func GetPortsStatus(ip string, communit string) (ports map[int]*model.Port, err error) {
	oids := []string{IfIndex, IfAdminStatus, IfOperStatus}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return ports, err
	}
	ports = make(map[int]*model.Port)
	for _, row := range tableRows {
		port := new(model.Port)
		port.Index = GetSnmpInt(row[IfIndex])
		port.AdminStatus = GetSnmpInt(row[IfAdminStatus])
		port.OperStatus = GetSnmpInt(row[IfOperStatus])
		ports[port.Index] = port
	}
	return
}

//GetPortInformation device,communit,index string
func GetPortInformation(ip, communit, index string) (port *model.Port, err error) {
	s := NewSNMP(ip, communit)
	err = s.Connect()
	if err != nil {
		return
	}
	defer s.Conn.Close()

	oids := []string{IfDescr, IfAdminStatus, IfOperStatus, IfSpeed}
	uc.Each(oids, func(r string, i int) {
		oids[i] = oids[i] + "." + index
	})
	result, err := s.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err != nil {
		return
	}
	port = new(model.Port)
	port.Index, _ = strconv.Atoi(index)
	for _, v := range result.Variables {
		switch v.Name {
		case oids[0]:
			port.Name = GetSnmpString(v)
		case oids[1]:
			port.AdminStatus = GetSnmpInt(v)
		case oids[2]:
			port.OperStatus = GetSnmpInt(v)
		case oids[3]:
			port.Speed = GetSnmpInt64(v)
		}
	}
	return

}
