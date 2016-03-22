package main

import (
	"io/ioutil"
	"fmt"
	"net"
	"strconv"
	"flag"
	"net/http"
)

func getFreePort(hostname string) int {
	var port_from, port_to int

	port_range, _ := ioutil.ReadFile("/proc/sys/net/ipv4/ip_local_port_range")
	fmt.Sscanf(string(port_range), "%d %d", &port_from, &port_to)

	for port := port_from; port <= port_to; port++ {
		ln, err := net.Listen("tcp", hostname + ":" + strconv.Itoa(port))
		if err == nil {
			ln.Close()
			return port
		}
	}
	return -1
}

func main() {

	var hostname string
	flag.StringVar(&hostname, "hostname", "", "Hostname to bind to")
	port := flag.Int("port", -1, "Port to listen to. If not set a random free port will be used")
	path := flag.String("path", ".", "Root folder to serve")
	flag.Parse()

	if *port == -1 {
		*port = getFreePort(hostname)
	}

	if hostname == "" {
		fmt.Println("Serving at: http://localhost:" + strconv.Itoa(*port))
	} else {
		fmt.Println("Serving at: http://" + hostname + ":" + strconv.Itoa(*port))
	}
	http.ListenAndServe(hostname + ":" + strconv.Itoa(*port), http.FileServer(http.Dir(*path)))

}
