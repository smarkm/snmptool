package snmp

import (
	"fmt"
	"strconv"

	"github.com/smarkm/snmptool/snmp/model"
	g "github.com/soniah/gosnmp"
)

//GetSystem get base system information
func GetSystem(ip string, communit string, ver g.SnmpVersion) (d model.Device, e error) {
	return GetDeviceInfo(ip, communit, ver)
}

//Ping snmp pring
func Ping(ip string, communit string, ver g.SnmpVersion) (d model.Device, e error) {
	return GetSystem(ip, communit, ver)
}

//GetHostName RT
func GetHostName(ip string, community string, ver g.SnmpVersion) (host string, err error) {
	device, err := GetSystem(ip, community, ver)
	if err != nil {
		return "", err
	}
	return device.Name, nil
}

//GetDeviceInfo get base system information
func GetDeviceInfo(ip string, communit string, ver g.SnmpVersion) (d model.Device, err error) {
	s := NewSNMP(ip, communit, ver)
	err = s.Connect()
	if err != nil {
		return
	}
	defer s.Conn.Close()

	oids := []string{SysObjectID, SysDescr, SysUpTime, SysContact, SysName, SysLocation, SysServices}
	result, err := s.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err != nil {
		return
	}
	for _, v := range result.Variables {
		switch v.Name {
		case SysObjectID:
			d.OId = v.Value.(string)
		case SysDescr:
			bytes := v.Value.([]byte)
			d.Desc = string(bytes)
		case SysUpTime:
			d.UpTime = fmt.Sprint(v.Value)
		case SysContact:
			bytes := v.Value.([]byte)
			d.Contract = string(bytes)
		case SysName:
			bytes := v.Value.([]byte)
			d.Name = string(bytes)
		case SysLocation:
			bytes := v.Value.([]byte)
			d.Location = string(bytes)
		case SysServices:
			bytes := v.Value.(int)
			d.Services = strconv.Itoa(bytes)
		}
	}
	d.IP = ip
	return
}
