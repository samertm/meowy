package main

import (
	"flag"
	"fmt"

	"github.com/samertm/meowy/server"
)

func main() {
	host := flag.String("host", "localhost", "sets the host name.")
	port := flag.String("port", "5849", "sets the port.")
	flag.Parse()
	ip := *host + ":" + *port
	fmt.Println("listening on", ip)
	server.ListenAndServe(ip)
}
