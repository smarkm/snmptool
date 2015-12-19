package snmp

//LLDP
const (
	//LLdpLocPortSubtype = [...]string{""}
	LldpLocPortNum       = ".1.0.8802.1.1.2.1.3.7.1.1"
	LldpLocPortIDSubtype = ".1.0.8802.1.1.2.1.3.7.1.2"
	LldpLocPortID        = ".1.0.8802.1.1.2.1.3.7.1.3"
	LldpLocPortDesc      = ".1.0.8802.1.1.2.1.3.7.1.4"

	LLdpRemTimeMark         = ".1.0.8802.1.1.2.1.4.1.1.1"
	LLdpRemLocalPortNum     = ".1.0.8802.1.1.2.1.4.1.1.2"
	LLdpRemIndex            = ".1.0.8802.1.1.2.1.4.1.1.3"
	LLdpRemChassisIDSubtype = ".1.0.8802.1.1.2.1.4.1.1.4"
	LLdpRemChassisID        = ".1.0.8802.1.1.2.1.4.1.1.5"
	LLdpRemPortIDSubtype    = ".1.0.8802.1.1.2.1.4.1.1.6"
	LLdpRemPortID           = ".1.0.8802.1.1.2.1.4.1.1.7"
	LLdpRemPortDesc         = ".1.0.8802.1.1.2.1.4.1.1.8"
	LLdpRemSysName          = ".1.0.8802.1.1.2.1.4.1.1.9"
	LLdpRemSysDesc          = ".1.0.8802.1.1.2.1.4.1.1.10"
	LLdpRemSysCapSupported  = ".1.0.8802.1.1.2.1.4.1.1.11"
	LLdpRemSysCapEnabled    = ".1.0.8802.1.1.2.1.4.1.1.12"
)

//LLdpLoc lldap
type LLdpLoc struct {
	LldpLocPortNum       string
	LldpLocPortIDSubtype string
	LldpLocPortID        string
	LldpLocPortDesc      string
}

//LLdpRem lldprem
type LLdpRem struct {
	LLdpRemLocalPortNum string
	LLdpRemSysName      string
	LLdpRemPortID       string
	LLdpRemPortDesc     string
}

//GetLLdpLocalTable get loclTable
func GetLLdpLocalTable(ip string, communit string) (lldpls []*LLdpLoc, err error) {
	oids := []string{LldpLocPortNum, LldpLocPortID, LldpLocPortDesc, LldpLocPortIDSubtype}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		lldpl := new(LLdpLoc)
		lldpl.LldpLocPortNum = GetSnmpString(row[LldpLocPortNum])
		lldpl.LldpLocPortID = GetSnmpString(row[LldpLocPortID])
		lldpl.LldpLocPortDesc = GetSnmpString(row[LldpLocPortDesc])
		lldpl.LldpLocPortIDSubtype = GetSnmpString(row[LldpLocPortIDSubtype])
		lldpls = append(lldpls, lldpl)
	}
	return
}

//GetLLdpRemTable get remote table
func GetLLdpRemTable(ip string, communit string) (lldpls []*LLdpRem, err error) {
	oids := []string{LLdpRemSysName, LLdpRemPortID, LLdpRemPortDesc}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return
	}
	for index, row := range tableRows {
		lldpl := new(LLdpRem)
		lldpl.LLdpRemLocalPortNum = index
		lldpl.LLdpRemSysName = GetSnmpString(row[LLdpRemSysName])
		lldpl.LLdpRemPortID = GetSnmpString(row[LLdpRemPortID])
		lldpl.LLdpRemPortDesc = GetSnmpString(row[LLdpRemPortDesc])
		lldpls = append(lldpls, lldpl)
	}
	return
}
