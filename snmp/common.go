package snmp

import (
	"fmt"
	"time"

	g "github.com/soniah/gosnmp"
)

//System
const (
	SysObjectID = ".1.3.6.1.2.1.1.2.0"
	SysDescr    = ".1.3.6.1.2.1.1.1.0"
	SysUpTime   = ".1.3.6.1.2.1.1.3.0"
	SysContact  = ".1.3.6.1.2.1.1.4.0"
	SysName     = ".1.3.6.1.2.1.1.5.0"
	SysLocation = ".1.3.6.1.2.1.1.6.0"
	SysServices = ".1.3.6.1.2.1.1.7.0"
)

//MSG
const (
	InvalidValue = "invalid value"
)

func init() {
	OIDs["sysObjectID"] = SysObjectID
	OIDs["sysDescr"] = SysDescr
	OIDs["sysUpTime"] = SysUpTime
	OIDs["sysContact"] = SysContact
	OIDs["sysName"] = SysName
	OIDs["sysLocation"] = SysLocation
	OIDs["sysServices"] = SysServices

}

//NewSNMP get a snmp instance
func NewSNMP(target string, community string) (s *g.GoSNMP) {
	s = &g.GoSNMP{
		Port:      161,
		Community: community,
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   3,
		MaxOids:   g.MaxOids,
		Target:    target,
	}
	return
}

//GetTable implement by BulkWalkAll
func GetTable(ip string, communit string, oids []string) (tableRows map[string](map[string]g.SnmpPDU), e error) {
	s := NewSNMP(ip, communit)
	err := s.Connect()
	if err != nil {
		return
	}
	defer s.Conn.Close()

	tableRows = make(map[string](map[string]g.SnmpPDU))
	for _, oid := range oids {
		rows, e := s.BulkWalkAll(oid)
		if e != nil {
			return tableRows, e
		}
		for _, row := range rows {
			index := row.Name[(len(oid) + 1):]
			r := tableRows[index]
			if r == nil {
				r = make(map[string]g.SnmpPDU)
			}
			r[oid] = row
			tableRows[index] = r
		}
	}

	return
}

//GetOne ip,comunity,oid string
func GetOne(ip, community, oid string) (rs g.SnmpPDU, err error) {
	s := NewSNMP(ip, community)
	err = s.Connect()
	if err != nil {
		return
	}
	defer s.Conn.Close()

	oids := []string{oid}
	result, err := s.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err != nil {
		return
	}
	for _, v := range result.Variables {
		switch v.Name {
		case oid:
			rs = v
		}
	}
	return
}

//GetSnmpString convert result 2 string
func GetSnmpString(p g.SnmpPDU) (v string) {
	switch p.Type {
	case g.OctetString:
		bytes := p.Value.([]byte)
		v = string(bytes)
	case g.ObjectIdentifier:
		v = p.Value.(string)
	default:
		v = fmt.Sprint(p.Value)
	}
	return
}

//GetSnmpMacString get mac address
func GetSnmpMacString(p g.SnmpPDU) (v string) {
	switch p.Type {
	case g.OctetString:
		bytes := p.Value.([]byte)
		split := ""
		for _, b := range bytes {
			if v != "" {
				split = ":"
			}
			v += fmt.Sprintf(split+"%x", b)
		}
	default:
		v = fmt.Sprint(p.Value)
	}
	return
}

//GetSnmpInt convert result 2 itn
func GetSnmpInt(p g.SnmpPDU) (v int) {
	switch p.Type {
	case g.Integer:
		v = p.Value.(int)
	case g.Counter32:
		fmt.Println(p.Type)
	}
	return
}

//GetSnmpInt64 get counter64
func GetSnmpInt64(p g.SnmpPDU) (v int64) {
	switch p.Type {
	case g.Integer:
		tmp := p.Value.(int)
		v = int64(tmp)
	case g.Counter32:
		fmt.Println(p.Type)
	case g.Gauge32:
		{
			t := p.Value.(uint)
			v = int64(t)
		}

	}
	return
}
