package main

import (
	"flag"
	"fmt"
)

var (
	addr         string
	dataDir      string
	debugEnabled bool
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1:8085", "ipxe reachable local address")
	flag.StringVar(&dataDir, "data-dir", "", "directory containing boot client definitions")
	flag.BoolVar(&debugEnabled, "debug", false, "increases logging verbosity level")
	flag.Parse()

	// handle the debug thingy

	fmt.Println("vim-go")
}
