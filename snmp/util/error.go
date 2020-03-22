package util

import (
	"fmt"
	"strings"
)

//HandleError RT
func HandleError(err error) {
	msg := err.Error()
	fmt.Println(msg)
	if strings.Contains(msg, "timeout") {
		fmt.Println("Report to: https://github.com/smarkm/snmptool/issues")
	}
}
