# csirtgsdk-go
Go Implementation of the CSIRTGSDK https://csirtg.io

# Getting Started

```bash
$ go run main.go 
# id,indicator,itype,portlist,firsttime,count,protocol,application,asn,cc,lasttime,description,provider 
12913172,178.46.48.38,ipv4,23,2018-12-01 21:31:40 UTC,1,6,,12389.0,RU ,2018-12-01 21:31:40 UTC,sourced from firewall logs (incomming  tcp  syn  blocked),
12913206,112.68.55.26,ipv4,23,2018-12-01 21:40:51 UTC,1,6,,17511.0,JP ,2018-12-01 21:40:51 UTC,sourced from firewall logs (incomming  tcp  syn  blocked),
12913212,79.133.98.142,ipv4,3389,2018-12-01 21:41:38 UTC,1,6,,57311.0,RU ,2018-12-01 21:41:38 UTC,sourced from firewall logs (incomming  tcp  syn  blocked),
..

$ go run main.go -user wes -feed darknet
...

$ go build -o csirtg main.go
```