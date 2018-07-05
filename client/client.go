package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

/*
Client - proxy client that connects to Proxy service and gets new IP for proxy
Otherwise - inits default http.Client
*/
type Client struct {
	key                string
	proxyServerAddress string // HTTP proxy server address
	client             *http.Client
}

// New - Client constructor
func New(proxyServerAddr string, key string) *Client {
	return &Client{
		key:                key,
		proxyServerAddress: proxyServerAddr,
		client:             &http.Client{},
	}
}

// ReNew - getting new IP address
func (c *Client) ReNew() {
	proxyAddr, err := getProxyAddress(c.proxyServerAddress, c.key)
	if err != nil {
		log.Println("[ERROR] Getting proxy IP: ", err)
	}
	if err == nil {
		// Setting Proxy
		if proxyURL, err := url.Parse(proxyAddr); err == nil {
			c.client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
		}
	}
}

// Do - HTTP request doer
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func getProxyAddress(addr, key string) (ip string, err error) {
	proxyServerAddr := addr + "/" + key
	if !strings.Contains(proxyServerAddr, "http") {
		proxyServerAddr = "http://" + proxyServerAddr
	}
	res, err := http.Get(proxyServerAddr)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}
	ip = string(b)
	return
}
