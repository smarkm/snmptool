package snmp

import (
	"testing"
	"time"
)

func TestTrapReciver_OnNewTrap(t *testing.T) {
	r := &TrapReciver{Address: "172.16.2.2"}
	defer r.Close()
	r.Start()
	time.Sleep(30 * time.Second)
}
