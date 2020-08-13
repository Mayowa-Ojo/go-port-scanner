package mod

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	hostname = "localhost"
	timeout  = 60 * time.Second
)

// ScanPorts - checks all network ports and reports the ones in use
func ScanPorts(start int, end int, protocol string) {
	for i := start; i < end; i++ {
		isOpen := ScanPort(protocol, hostname, i, timeout)

		if isOpen {
			log.Printf("port '%d' is in use", i)
		}
	}
}

// ScanPort - checks if a port is currently in use
func ScanPort(protocol string, hostname string, port int, timeout time.Duration) bool {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, addr, timeout)

	if err != nil {
		return false
	}

	if conn != nil {
		defer conn.Close()
		return true
	}

	return false
}

// {int} p - number of processes
// {int} r - number of goroutines

// LaunchGoroutines - starts a number of goroutines indicated by <r> to run concurrent processes
func LaunchGoroutines(p, r int, f func(int, int, string), protocol string, wg *sync.WaitGroup) {
	for i := r; i < p; i += r {
		wg.Add(1)

		go func(nextIndex int) {
			f(nextIndex-100, nextIndex, protocol)
			wg.Done()
		}(i)

	}
}
