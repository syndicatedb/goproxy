package main

import (
	"flag"
	"io/ioutil"
	"strings"

	"github.com/syndicatedb/goproxy"
)

var (
	dataFile = flag.String("data-file", "./ip.txt", "Path to file which contains list of IP-addresses")
)

func main() {
	flag.Parse()

	ipx := loadIP(*dataFile)
	srv := goproxy.New(":8081", ipx)
	srv.Start()
}

func loadIP(file string) (ipx []string) {
	var b []byte
	var err error
	if b, err = ioutil.ReadFile(file); err != nil {
		panic("Error reading file")
	}
	ipx = strings.Split(string(b), "\n")
	return
}
