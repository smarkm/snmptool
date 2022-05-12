package snmp

import g "github.com/soniah/gosnmp"

//BGP oids
const (
	BgpLocalAS = ".3.6.1.2.1.15.2.0"

	BgpPeerLocalAddrOid  = ".1.3.6.1.2.1.15.3.1.5"
	BgpPeerLocalPortOid  = ".1.3.6.1.2.1.15.3.1.6"
	BgpPeerRemoteAddrOid = ".1.3.6.1.2.1.15.3.1.7"
	BgpPeerRemotePortOid = ".1.3.6.1.2.1.15.3.1.8"
	BgpPeerRemoteASOid   = ".1.3.6.1.2.1.15.3.1.9"
	BgpPeerStateOid      = ".1.3.6.1.2.1.15.3.1.2"
)

//Var
var (
	PeerState = []string{"idle(1)", "connect(2)", "active(3)", "opensent(4)", "openconfirm(5)", "established(6)"}
)

func init() {
	OIDs["bgpLocalAS"] = BgpLocalAS
	OIDs["bgpPeerLocalAddr"] = BgpPeerLocalAddrOid
	OIDs["bgpPeerLocalPort"] = BgpPeerLocalPortOid
	OIDs["bgpPeerRemoteAddr"] = BgpPeerRemoteAddrOid
	OIDs["bgpPeerRemotePort"] = BgpPeerRemotePortOid
	OIDs["bgpPeerRemoteAS"] = BgpPeerRemoteASOid
	OIDs["bgpPeerStateOid"] = BgpPeerStateOid
}

//BGPPeer snmp struct
type BGPPeer struct {
	LocalAddr  string
	LocalPort  int
	RemoteAddr string
	RemotePort int
	RemoteAS   int
	PeerState  int
}

//PeerStateStr readable state
func (p *BGPPeer) PeerStateStr() string {
	if p.PeerState > 0 {
		return PeerState[p.PeerState-1]
	}
	return InvalidValue
}

//GetLocalAS get local as
func GetLocalAS(s g.GoSNMP) (as int, err error) {
	pdu, err := GetOne(s, BgpLocalAS)
	if err != nil {
		return
	}
	as = GetSnmpInt(pdu)
	return
}

//GetBGPPeerTable get Nbr table
func GetBGPPeerTable(s g.GoSNMP) (rs []*BGPPeer, err error) {
	oids := []string{BgpPeerLocalAddrOid, BgpPeerLocalPortOid, BgpPeerRemoteAddrOid,
		BgpPeerRemotePortOid, BgpPeerRemoteASOid, BgpPeerStateOid}
	tableRows, err := GetTable(s, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		item := new(BGPPeer)
		item.LocalAddr = GetSnmpString(row[BgpPeerLocalAddrOid])
		item.LocalPort = GetSnmpInt(row[BgpPeerLocalPortOid])
		item.RemoteAddr = GetSnmpString(row[BgpPeerRemoteAddrOid])
		item.RemotePort = GetSnmpInt(row[BgpPeerRemotePortOid])
		item.RemoteAS = GetSnmpInt(row[BgpPeerRemoteASOid])
		item.PeerState = GetSnmpInt(row[BgpPeerStateOid])
		rs = append(rs, item)
	}
	return
}
