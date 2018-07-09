package goproxy

import "sync"

type balancer struct {
	list  []IP
	state map[string]int
	sync.Mutex
}

// IP - struct that contains IP address
type IP struct {
	address string
}

// newBalancer - balancer constructor
func newBalancer(ipx []string) *balancer {
	return &balancer{
		list:  populateList(ipx),
		state: make(map[string]int),
	}
}

func (b *balancer) issue(key string) (ip string, err error) {
	b.Lock()
	defer b.Unlock()

	var pos int // last position in slice for this key
	pos = b.state[key]
	index := len(b.list) - 1
	if pos > index {
		pos = 0
	}
	ip = b.list[pos].address
	pos++
	b.state[key] = pos
	return
}

func populateList(ipx []string) (list []IP) {
	for _, ip := range ipx {
		// TODO: IP validation
		list = append(list, IP{
			address: ip,
		})
	}
	return
}
