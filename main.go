package main

import (
	"fmt"

	"github.com/samertm/meowy/server"
)

func main() {
	ip := "localhost:5849"
	fmt.Println("listening on", ip)
	server.ListenAndServe(ip)
}
