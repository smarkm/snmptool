package model

import (
	"strconv"
	"time"
)

//common
var (
	AdminStatus = []string{"up(1)", "down(2)", "testing(3)"}
	OperStatus  = []string{"up(1)", "down(2)", "testing(3)", "unknown(4)", "dormant(5)", "notPresent(6)", "lowerLayerDown(7)"}
)

//Port physical and logic
type Port struct {
	Index        int       `json:"index"`
	ID           string    `json:"id"`
	IP           string    `json:"ip"`
	Mac          string    `json:"mac"`
	Name         string    `json:"name"`
	Descr        string    `json:"descr"`
	UsefullSpeed int64     `json:"realSpeed"` //bits
	Speed        int64     `json:"speed"`     //bits
	Mtu          int       `json:"mtu"`
	AdminStatus  int       `json:"admin"`
	OperStatus   int       `json:"oper"`
	CreateTime   time.Time `json:"createTime"`
	SyncTime     time.Time `json:"syncTime"`
	DeviceIP     string    `json:"deviceIP"`
	DeviceName   string    `json:"deviceName"`
	Removed      bool      `json:"removed"`
	RemovedTime  time.Time `json:"removedTime"`
	Monitor      bool      `json:"monitor"`
}

//RealSpeed get Real UsefullSpeed
func (p *Port) RealSpeed() (s int64) {
	s = p.Speed
	if p.UsefullSpeed > 0 {
		s = p.UsefullSpeed
	}
	return
}

//BundleEther logic port
type BundleEther struct {
	Port
	SubIndexes []string
}

//AdminStr
func (p *Port) AdminStr() (v string) {
	if p.AdminStatus > len(AdminStatus) {
		v = strconv.Itoa(p.AdminStatus)
	} else {
		v = AdminStatus[p.AdminStatus-1]
	}
	return
}

//AdminStr
func (p *Port) OperStr() (v string) {
	if p.OperStatus > len(OperStatus) {
		v = strconv.Itoa(p.OperStatus)
	} else {
		v = OperStatus[p.OperStatus-1]
	}
	return
}
