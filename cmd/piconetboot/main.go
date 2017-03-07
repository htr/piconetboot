package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	const defaultScript = `#!ipxe
dhcp
params
set idx:int32 0
:loop isset ${net${idx}/mac} || goto loop_done
echo net${idx} is a ${net${idx}/chip} with MAC ${net${idx}/mac}
param net${idx}mac ${net${idx}/mac}
inc idx && goto loop
:loop_done
param uuid ${uuid}
param serial ${serial}
param asset ${asset}
chain http://10.0.15.1:8085/boot##params
`

	const bootScript = `#!ipxe
set base-url http://stable.release.core-os.net/amd64-usr/current
kernel ${base-url}/coreos_production_pxe.vmlinuz console=tty0 console=ttyS0 coreos.autologin=tty1 coreos.autologin=ttyS0
initrd ${base-url}/coreos_production_pxe_image.cpio.gz
boot
`

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(defaultScript))
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}

	log.Fatal(srv.ListenAndServe())

	fmt.Println("vim-go")
}
