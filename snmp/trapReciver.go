package snmp

import (
	"log"
	"net"

	"github.com/soniah/gosnmp"
)

var (
	recivers = make(map[string]*TrapReciver)
)

//TrapReciver RT
type TrapReciver struct {
	Address     string
	i           *gosnmp.TrapListener
	TrapHandler func(s *gosnmp.SnmpPacket, u *net.UDPAddr)
}

//Start start an Reciver instance
func (r *TrapReciver) Start() {

	recivers[r.Address] = r
	log.Println("Trap Receiver " + r.Address + " running ...")
	r.i = gosnmp.NewTrapListener()
	if r.TrapHandler != nil {
		r.i.OnNewTrap = r.TrapHandler
	}
	r.i.Listen(r.Address)
}

//Close instance
func (r *TrapReciver) Close() {
	address := r.Address
	r.i.Close()
	delete(recivers, address)
}

//registDeviceTrapHandler device level handler
func registDeviceTrapHandler(s *gosnmp.SnmpPacket, u *net.UDPAddr) {
	tmp := make(map[string]gosnmp.SnmpPDU)
	for _, pdu := range s.Variables {
		tmp[pdu.Name] = pdu
	}
	trapOid := GetSnmpString(tmp[TrapOID])
	switch trapOid {
	case WarmStart:
		log.Println("warmstart")
	default:
		log.Println("unknow trap")
	}

}
