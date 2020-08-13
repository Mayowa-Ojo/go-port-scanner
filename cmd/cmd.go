package cmd

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/Mayowa-Ojo/go-port-scanner/mod"
)

var port int
var wg sync.WaitGroup

// Execute - CLI Entry point
func Execute() {
	flag.Parse()

	if port == 0 {
		// scan all ports
		mod.LaunchGoroutines(65000, 100, mod.ScanPorts, &wg)

		wg.Wait()
		return
	}

	isClosed := mod.ScanPort("tcp", "localhost", port, 60*time.Second)

	if isClosed {
		log.Printf("port '%d' is in use", port)
	} else {
		log.Printf("port '%d' is idle", port)
	}

}

func init() {
	flag.IntVar(&port, "port", 0, "port number to be scanned")
}
