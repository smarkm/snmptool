package snmp

import g "github.com/soniah/gosnmp"

//Const
const (
	OIDHrProcessorLoad = ".1.3.6.1.2.1.25.3.3.1.2"
)

//HrProcessorLoad process load wap
type HrProcessorLoad struct {
	Load int
}

//GetHrProcessorLoad get loclTable
func GetHrProcessorLoad(s g.GoSNMP) (table []*HrProcessorLoad, err error) {
	oids := []string{OIDHrProcessorLoad}
	tableRows, err := GetTable(s, oids)
	if err != nil {
		return
	}
	for _, row := range tableRows {
		processor := new(HrProcessorLoad)
		processor.Load = GetSnmpInt(row[OIDHrProcessorLoad])
		table = append(table, processor)
	}
	return
}
