package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/htr/piconetboot/localfs"

	"github.com/gorilla/mux"
)

var (
	addr         string
	dataDir      string
	staticDir    string
	debugEnabled bool
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1:8085", "ipxe reachable local address")
	flag.StringVar(&dataDir, "data-dir", "", "directory containing boot client definitions")
	flag.StringVar(&staticDir, "static-dir", "", "directory containing static files")
	flag.BoolVar(&debugEnabled, "debug", false, "increases logging verbosity level")
	flag.Parse()

	if debugEnabled {
		log.SetLevel(log.DebugLevel)
	}

	clientStore, err := localfs.NewStore(dataDir)

	if err != nil {
		log.Error("unable to create localfs store:", err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(genDefaultIpxeScript()))
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() != nil {
			w.WriteHeader(400)
			return
		}
		log.WithField("path", r.URL.Path).WithField("client-data", r.Form).WithField("headers", r.Header).Debug("Boot")

		client, err := clientStore.Find(r.Form)
		if err != nil {
			log.WithError(err).Info("unable to find a matching client")
			w.Write([]byte("#!ipxe\necho unable to find a matching client\nshell\n"))
			return
		}

		w.Write([]byte(ipxeScriptPreamble() + client.BootScript()))
	}).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	log.WithField("listen-addr", addr).Info("starting http server")

	srv := &http.Server{
		Handler: requestLogger{r},
		Addr:    addr,
	}

	log.Fatal(srv.ListenAndServe())

	fmt.Println("vim-go")
}

func genDefaultIpxeScript() string {
	return ipxeScriptPreamble() + `

dhcp
params
set idx:int32 0
:loop isset ${net${idx}/mac} || goto loop_done
echo net${idx} is a ${net${idx}/chip} with MAC ${net${idx}/mac}
param net${idx}mac ${net${idx}/mac}
param net${idx}bustype ${net${idx}/bustype}
param net${idx}busid ${net${idx}/busid}
param net${idx}chip ${net${idx}/chip}
param net${idx}busloc ${net${idx}/busloc}

inc idx && goto loop
:loop_done
param uuid ${uuid}
param serial ${serial}
param asset ${asset}

chain http://${bootserver}/##params
`
}

func ipxeScriptPreamble() string {
	localAddress := addr

	return fmt.Sprintf("#!ipxe\nset bootserver %s\n", localAddress)
}

type requestLogger struct {
	handler http.Handler
}

func (h requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.WithField("url", r.URL).WithField("remote-address", r.RemoteAddr).Info(r.Method)
	h.handler.ServeHTTP(w, r)
}
