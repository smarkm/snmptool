package cmd

import (
	"errors"
	"strconv"
)

func ParseIPAndCommunity(args []string, expect int) (ip, community string, err error) {
	ln := len(args)
	if ln < expect {
		err = errors.New("Error: Want at least " + strconv.Itoa(expect) + " args, now only " + strconv.Itoa(ln))
	}
	community = "public"
	switch ln {
	case 1:
		ip = args[0]
	case 2:
		ip = args[0]
		community = args[1]
	}
	return
}

//i2S string
func i2S(n int) string {
	return strconv.Itoa(n)
}
