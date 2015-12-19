package util

import (
	"fmt"
	"testing"
)

func TestGetDeviceType(t *testing.T) {
	oid := ".1.3.6.1.4.1.8072.3.2.10"
	fmt.Println(GetDeviceType(oid),GetDeviceIcon(oid))
}
