package mod

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	network  = "tcp"
	hostname = "localhost"
	timeout  = 60 * time.Second
)

// ScanPorts - checks all network ports and reports the ones in use
func ScanPorts(start int, end int) {
	for i := start; i < end; i++ {
		isClosed := ScanPort(network, hostname, i, timeout)

		if isClosed {
			log.Printf("port '%d' is in use", i)
		}
	}
}

// ScanPort - checks if a port is currently in use
func ScanPort(network string, hostname string, port int, timeout time.Duration) bool {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(network, addr, timeout)

	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}

// {int} p - number of processes
// {int} r - number of goroutines

// LaunchGoroutines - starts a number of goroutines indicated by <r> to run concurrent processes
func LaunchGoroutines(p, r int, f func(int, int), wg *sync.WaitGroup) {
	for i := r; i < p; i += r {
		wg.Add(1)

		go func(nextIndex int) {
			f(nextIndex-100, nextIndex)
			wg.Done()
		}(i)

	}
}
