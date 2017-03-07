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

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

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
