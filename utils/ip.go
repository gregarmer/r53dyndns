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

func GetInterfaceIP(inet string) (string, error) {
	// TODO
	ip := "127.0.0.4"
	if !ValidIP(ip) {
		return "", errors.New(fmt.Sprintf("Error: %s is not a valid IP", ip))
	}
	return ip, nil
}

func ValidIP(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}
