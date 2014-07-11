package main

import (
	"flag"
	"fmt"

	"github.com/samertm/meowy/server"
)

func main() {
	host := flag.String("host", "localhost", "sets the host name.")
	port := flag.String("port", "5849", "sets the port.")
	prefix := flag.String("prefix", "", "sets prefix (for if meowy listens on a path that isn't \"/\"")
	flag.Parse()
	ip := *host + ":" + *port
	fmt.Println("listening on", ip)
	if *prefix != "" {
		fmt.Println("with prefix", *prefix)
		var front, back string
		if (*prefix)[0] != '/' {
			front = "/"
		}
		if (*prefix)[len(*prefix)-1] != '/' {
			back = "/"
		}
		*prefix = front + *prefix + back
	}
	server.ListenAndServe(ip, *prefix)
}
