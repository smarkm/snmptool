## Snmptool
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9e720e9e45f6456abfa18d27d0b8136c)](https://app.codacy.com/app/smarkm/snmptool?utm_source=github.com&utm_medium=referral&utm_content=smarkm/snmptool&utm_campaign=Badge_Grade_Dashboard)
[![ASL 2.0](https://img.shields.io/hexpm/l/plug.svg)](https://github.com/smarkm/snmptool/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/smarkm/snmptool.svg?branch=master)](https://travis-ci.org/smarkm/snmptool)
[![](http://shields.katacoda.com/katacoda/smark/count.svg)](https://www.katacoda.com/smark/scenarios/snmptool)
[![Go Report Card](https://goreportcard.com/badge/github.com/smarkm/snmptool)](https://goreportcard.com/report/github.com/smarkm/snmptool)

## Download
download the last version
* Linux  : [snmptool](https://github.com/smarkm/snmptool/releases/download/v0.0.3/snmp)
* Window : [snmptool.exe](https://github.com/smarkm/snmptool/releases/download/v0.0.3/snmp.exe)

## Usage
![Show case](demo-v0.0.2.gif)
```
Simple snmp tool

Usage:
  snmptool [command]

Available Commands:
  bgp         Show BGP brief information
  cdp         Show CDP brief infromation
  get         Execute SNMP GET function
  help        Help about any command
  iftable     Show Iftable brief information
  interface   Show interface biref information
  ipaddress   Show ip address table
  lldp        Show LLDP brief information
  oids        Show Name-OID mapping
  ospf        Show OSPF biref information
  storage     Show storage biref information
  sys         Show system brief information
  trap        Start trap receiver
  version     Show version
  walk        Execute SNMP WALK

Flags:
  -c, --community string   community (default "public")
  -h, --help               help for snmptool
  -i, --ip string          target ip (default "127.0.0.1")

Use "snmptool [command] --help" for more information about a command.
```

* you can [play online](https://www.katacoda.com/smark/scenarios/snmptool)

## Road map


