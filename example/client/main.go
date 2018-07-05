package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/syndicatedb/goproxy/client"
)

func main() {
	c := client.New(":8081", "kucoin") // init proxy client
	c.ReNew()                          // getting shining new IP

	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := c.Do(req) // using as usual
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}
