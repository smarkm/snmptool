package snmp

const (
	BgpPeerLocalAddrOid  = ".1.3.6.1.2.1.15.3.1.5"
	BgpPeerLocalPortOid  = ".1.3.6.1.2.1.15.3.1.6"
	BgpPeerRemoteAddrOid = ".1.3.6.1.2.1.15.3.1.7"
	BgpPeerRemotePortOid = ".1.3.6.1.2.1.15.3.1.8"
	BgpPeerRemoteASOid   = ".1.3.6.1.2.1.15.3.1.9"
	BgpPeerStateOid      = ".1.3.6.1.2.1.15.3.1.2"
)

var (
	PeerState = []string{"idle(1)", "connect(2)", "active(3)", "opensent(4)", "openconfirm(5)", "established(6)"}
)

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

//GetOspfNbrTable get Nbr table
func GetBGPPeerTable(ip string, communit string) (rs []*BGPPeer, err error) {
	oids := []string{BgpPeerLocalAddrOid, BgpPeerLocalPortOid, BgpPeerRemoteAddrOid,
		BgpPeerRemotePortOid, BgpPeerRemoteASOid, BgpPeerStateOid}
	tableRows, err := GetTable(ip, communit, oids)
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
