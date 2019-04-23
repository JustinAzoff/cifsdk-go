# cifsdk-go
Go Implementation of the CIF SDK https://github.com/csirtgadgets/verbose-robot/

# Getting Started

```bash
$ export CIF_TOKEN=123123123

$ go run cmd/cifsdk-go/main.go  -feed ipv4 -endpoint http://192.168.99.100:5000
77.247.109.205,ipv4,,,3,209299,vitox telecom,identified as sending recursive dns queries to a remote host,dataplane.org
173.249.14.216,ipv4,,,2,51167,contabo gmbh,identified as sending recursive dns queries to a remote host,dataplane.org
81.171.71.127,ipv4,,,2,33438,"highwinds network group, inc.",identified as sending recursive dns queries to a remote host,dataplane.org
81.171.71.128,ipv4,,,2,33438,"highwinds network group, inc.",identified as sending recursive dns queries to a remote host,dataplane.org
...
