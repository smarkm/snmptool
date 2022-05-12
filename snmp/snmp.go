package snmp

import (
	"fmt"
	"strconv"

	"github.com/smarkm/snmptool/snmp/model"
	g "github.com/soniah/gosnmp"
)

//GetSystem get base system information
func GetSystem(s g.GoSNMP) (d model.Device, e error) {
	return GetDeviceInfo(s)
}

//Ping snmp pring
func Ping(s g.GoSNMP) (d model.Device, e error) {
	return GetSystem(s)
}

//GetHostName RT
func GetHostName(s g.GoSNMP) (host string, err error) {
	device, err := GetSystem(s)
	if err != nil {
		return "", err
	}
	return device.Name, nil
}

//GetDeviceInfo get base system information
func GetDeviceInfo(s g.GoSNMP) (d model.Device, err error) {
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
	d.IP = s.Target
	return
}
