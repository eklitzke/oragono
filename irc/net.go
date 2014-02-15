package irc

import (
	"log"
	"net"
	"strings"
)

func IPString(addr net.Addr) string {
	addrStr := addr.String()
	ipaddr, _, err := net.SplitHostPort(addrStr)
	if err != nil {
		return addrStr
	}
	return ipaddr
}

func AddrLookupHostname(addr net.Addr) string {
	return LookupHostname(IPString(addr))
}

func LookupHostname(addr string) string {
	names, err := net.LookupAddr(addr)
	if err != nil {
		return addr
	}

	hostname := strings.TrimSuffix(names[0], ".")
	return hostname
}
