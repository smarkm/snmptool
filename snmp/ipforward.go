package snmp

//IPforward table itme ,.1.3.6.1.2.1.4.24.2
const (
	IPForwardDest    = ".1.3.6.1.2.1.4.24.2.1.1"
	IPForwardMask    = ".1.3.6.1.2.1.4.24.2.1.2"
	IPForwardNextHop = ".1.3.6.1.2.1.4.24.2.1.4"
	IPForwardIfIndex = ".1.3.6.1.2.1.4.24.2.1.5"
)

//IPForwardItem from if table
type IPForwardItem struct {
	Dest    string `json:"dest"`
	Mask    string `json:"mask"`
	NextHop string `json:"nexthop"`
	IfIndex int    `json:"ifIndex"`
}

//GetIPForwardTable get loclTable
func GetIPForwardTable(ip string, communit string) (items []*IPForwardItem, err error) {
	oids := []string{IPForwardDest, IPForwardMask, IPForwardNextHop, IPForwardIfIndex}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		item := new(IPForwardItem)
		item.Dest = GetSnmpString(row[IPForwardDest])
		item.Mask = GetSnmpString(row[IPForwardMask])
		item.NextHop = GetSnmpString(row[IPForwardNextHop])
		item.IfIndex = GetSnmpInt(row[IPForwardIfIndex])
		items = append(items, item)
	}
	return
}
