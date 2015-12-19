package util

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var deviceTypes = make(map[string]string)
var deviceIcons = make(map[string]string)

func init() {
	deviceTypes[".1.3.6.1.4.1.311.1.1.3.1.1"] = "Windows"
	csvFile, err := os.Open("config/dt.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	defer csvFile.Close()
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		deviceTypes[line[0]] = line[1]
		deviceIcons[line[0]] = line[2]
	}

}

//GetDeviceType get device type base oid
func GetDeviceType(sysObjectID string) (t string) {
	t, ok := deviceTypes[sysObjectID]
	if false == ok {
		t = Unknow
	}
	return
}
func GetDeviceIcon(sysObjectID string) (t string) {
	t, ok := deviceIcons[sysObjectID]
	if false == ok {
		t = ""
	}
	return
}
