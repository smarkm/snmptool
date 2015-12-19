package snmp

const (
	OID_HrProcessorLoad = ".1.3.6.1.2.1.25.3.3.1.2"
)

type HrProcessorLoad struct {
	Load int
}

//GetLLdpLocalTable get loclTable
func GetHrProcessorLoad(ip string, communit string) (table []*HrProcessorLoad, err error) {
	oids := []string{OID_HrProcessorLoad}
	tableRows, err := GetTable(ip, communit, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		processor := new(HrProcessorLoad)
		processor.Load = GetSnmpInt(row[OID_HrProcessorLoad])
		table = append(table, processor)
	}
	return
}
