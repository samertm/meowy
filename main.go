// meowy
// Copyright (C) 2014 Samer Masterson

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Contact the author by email: samer@samertm.com

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
		var front, back bool
		if (*prefix)[0] != '/' {
			front = true
		}
		if (*prefix)[len(*prefix)-1] == '/' {
			back = true
		}
		if front {
			*prefix = "/" + *prefix
		}
		if back {
			*prefix = (*prefix)[0:len(*prefix)-1]
		}
		fmt.Println("with prefix", *prefix)
	}
	server.ListenAndServe(ip, *prefix)
}
