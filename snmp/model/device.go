package model

//Device l1,l2,l3 device
type Device struct {
	IP        string `gorm:"primary_key"`
	Type      string `json:"type"`
	Vendor    string `json:"vendor"`
	OId       string `json:"oid"`
	Desc      string `json:"description"`
	Status    string `json:"status"`
	Location  string
	Contract  string
	Services  string
	Name      string `json:"name"`
	Community string `json:"community"`
}
