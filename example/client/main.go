package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/syndicatedb/goproxy/proxy"
)

func main() {
	p := proxy.New(":8081") // init proxy Fabric
	c := p.NewClient("kucoin")

	req, err := http.NewRequest("GET", "https://google.com", nil)
	req.Header.Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		panic(err)
	}

	resp, err := c.Do(req) // using as usual
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}
