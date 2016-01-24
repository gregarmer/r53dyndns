package main

import (
	// "bytes"
	"flag"
	"github.com/gregarmer/r53dyndns/config"
	"github.com/gregarmer/r53dyndns/utils"
	"io/ioutil"
	"log"
	// "os"
	// "os/exec"
	// "path/filepath"
	"time"
)

const workingDir = "temp"

var configFile = flag.String("c", "~/.r53dyndns", "path to the config file")
var domain = flag.String("d", "", "domain to update")
var verbose = flag.Bool("v", false, "be verbose")

func main() {
	start_time := time.Now()

	flag.Parse()

	if *domain == "" {
		utils.Fatalf("Error: please supply the domain to update. Use -h for help.")
	}

	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}

	log.Printf("starting Route 53 DNS update for %s", *domain)

	conf := config.LoadConfig(*configFile)

	log.Print(conf.AwsAccessKey)

	log.Printf("done - took %s", time.Since(start_time))
}
