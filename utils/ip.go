package utils

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

func GetExternalIP() (string, error) {
	// dig +short myip.opendns.com @resolver1.opendns.com
	target := "myip.opendns.com."
	server := "resolver1.opendns.com:53"

	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(target, dns.TypeA)
	r, t, err := c.Exchange(&m, server)
	CheckErr(err)

	if len(r.Answer) < 1 {
		log.Fatal("No results")
	}

	firstRecord := r.Answer[0].(*dns.A)
	ip := fmt.Sprintf("%s", firstRecord.A)
	log.Printf("found external IP %s in %v", firstRecord.A, t)

	if !ValidIP(ip) {
		return "", errors.New(fmt.Sprintf("Error: %s is not a valid IP", ip))
	}

	return ip, nil
}

func GetInterfaceIP(requestedIface string) (string, error) {
	ipA := ""

	// try to get the requested interface
	iface, err := net.InterfaceByName(requestedIface)
	if err != nil {
		return "", err
	}

	// grab all addresses assigned to this interface
	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	// loop through addresses trying to find a v4 address (future v6 support)
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		// process IP address
		log.Printf("found ip address: %s", ip)

		// if it's IPv4 we'll use it
		if ip.To4() != nil {
			ipA = ip.String()
		}
	}

	if !ValidIP(ipA) {
		return "", errors.New(fmt.Sprintf("Error: %s is not a valid IP", ipA))
	}
	return ipA, nil
}

func ValidIP(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}
