package goproxy

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

func (r response) error(f string) string {
	if f == "json" {
		b, _ := json.Marshal(r)
		return string(b)
	}
	return r.Error.Error()
}
func (r response) data(f string) string {
	if f == "json" {
		b, _ := json.Marshal(r)
		return string(b)
	}
	return fmt.Sprintf("%v", r.Data)
}
