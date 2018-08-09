package goproxy

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	defaultAddress  = ":8080"
	contentTypeJSON = "application/json"
	typeJSON        = "json"
	typeTXT         = "txt"
)

// Server - serves proxy balancer
type Server struct {
	addr     string
	balancer *balancer
}

// New - Server constructor
func New(addr string, ipx []string) *Server {
	return &Server{
		addr:     addr,
		balancer: newBalancer(ipx),
	}
}

// Start - starting goproxy server
func (srv *Server) Start() {
	log.Println("Starting goproxy server on ", srv.addr)
	http.HandleFunc("/", srv.issue)
	log.Fatal(http.ListenAndServe(srv.addr, nil))
}

func (srv *Server) issue(w http.ResponseWriter, r *http.Request) {
	key := strings.Replace(r.URL.Path, "/", "", 1)
	ip, err := srv.balancer.issue(key)
	resp := response{
		Data:  ip,
		Error: err,
	}

	format := getFormat(r)
	if format == typeJSON {
		w.Header().Set("Content-Type", contentTypeJSON)
	}
	if err != nil {
		http.Error(w, resp.error(format), http.StatusInternalServerError)
	}
	fmt.Fprint(w, resp.data(format))
}

func getFormat(r *http.Request) (format string) {
	format = typeTXT
	if r.Header.Get("Accept") == contentTypeJSON {
		return typeJSON
	}
	q := r.URL.Query()
	if _, ok := q[typeJSON]; ok {
		return typeJSON
	}
	return
}
