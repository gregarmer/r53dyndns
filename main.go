package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/gregarmer/r53dyndns/config"
	"github.com/gregarmer/r53dyndns/dyndns"
	"github.com/gregarmer/r53dyndns/utils"
)

const workingDir = "temp"

var configFile = flag.String("c", "~/.r53dyndns", "path to the config file")
var domain = flag.String("d", "", "domain to update")
var staticIp = flag.String("s", "", "use this ip for the update")
var autoDetect = flag.Bool("a", false, "auto detect external IP")
var interfaceIp = flag.String("i", "", "use the IP from this interface for the update, eg: eth0")
var verbose = flag.Bool("v", false, "be verbose")

func main() {
	start_time := time.Now()
	var ip string

	flag.Parse()

	if *domain == "" {
		utils.Fatalf("Error: please supply the domain to update. Use -h for help.")
	}

	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}

	log.Printf("starting Route 53 DNS update for %s", *domain)

	conf := config.LoadConfig(*configFile)
	r53 := dyndns.Dyndns{Config: conf}

	switch {
	case *staticIp != "":
		ip = *staticIp
		if !utils.ValidIP(ip) {
			utils.Fatalf("Error: %s is not a valid IP", ip)
		}
	case *interfaceIp != "":
		var err error
		ip, err = utils.GetInterfaceIP(*interfaceIp)
		utils.CheckErr(err)
	case *autoDetect:
		var err error
		ip, err = utils.GetExternalIP()
		utils.CheckErr(err)
	default:
		utils.Fatalf("Error: please supply one of -s, -a, -i. Use -h for help.")
	}

	// update record
	// XXX: only if it changed from the last run
	r53.UpsertDomain(*domain, ip)

	log.Printf("done - took %s", time.Since(start_time))
}
