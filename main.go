package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/pichik/WebWatcher/src/auth"
	"github.com/pichik/WebWatcher/src/datacenter"
	"github.com/pichik/WebWatcher/src/datacenter/collector"
	"github.com/pichik/WebWatcher/src/datacenter/harvester"
	"github.com/pichik/WebWatcher/src/misc"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"golang.org/x/crypto/acme/autocert"
)

type Start struct {
	Token string
}
type End struct {
	Token string
}

func (s *Start) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	misc.RequestLog.Printf("%-17s %9s %s", "["+strings.Split(r.RemoteAddr, ":")[0]+"]", "["+r.Method+"]", r.RequestURI)
	SetHeaders(w)

	if r.Method == "OPTIONS" {
		misc.Custom404Handler(w)
		return
	}

	if collector.Baited(w, r) {
		w.WriteHeader(http.StatusOK)
		return
	}
	if !auth.CanAccess(r) {
		misc.Custom404Handler(w)
		return
	}

	next(w, r)
}

func (e *End) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if w.Header().Get("HasContent") == "true" {
		w.Header().Del("HasContent")
		w.WriteHeader(http.StatusOK)
		return
	}
	misc.Custom404Handler(w)

	next(w, r)
}

func main() {
	LoadFiles()

	r := mux.NewRouter()
	n := negroni.Classic()

	//Request handler order

	n.Use(&Start{})

	r.HandleFunc("/login", auth.Authenticate).Methods("GET")
	r.HandleFunc(datacenter.DeepCollectorPath, collector.DeepCollect)
	r.HandleFunc(auth.AdminPanel+"{id:[a-zA-Z0-9]{64}}", harvester.Extract).Methods("GET")
	r.HandleFunc(auth.AdminPanel+"{id:[a-zA-Z0-9]{64}}", harvester.Delete).Methods("DELETE")
	r.HandleFunc(auth.AdminPanel+"all", harvester.ExtractAll).Methods("GET")
	r.HandleFunc(auth.AdminPanel+"all", harvester.DeleteAll).Methods("DELETE")
	r.HandleFunc("/{id:"+misc.Config.CollectorPath+".*}", collector.RequestCollect)
	r.PathPrefix("/").Handler(http.HandlerFunc(Assets)).Methods("GET")

	n.UseHandler(r)

	n.Use(&End{})

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(misc.Config.Hostname, "www."+misc.Config.Hostname),
		Cache:      autocert.DirCache(misc.CertDir),
	}

	server := &http.Server{
		Addr:      ":https",
		Handler:   n,
		TLSConfig: certManager.TLSConfig(),
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(n))
	log.Fatal(server.ListenAndServeTLS("", ""))
}

func Assets(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		return
	}

	fs := http.FileServer(http.Dir(misc.AssetsDir))
	_, err := os.Stat(misc.AssetsDir + r.URL.Path)

	if err == nil {
		w.Header().Set("HasContent", "true")
		fs.ServeHTTP(w, r)
	}
}

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func LoadFiles() {
	misc.ImportLogs()
	misc.ImportTemplates()
	harvester.ImportTemplate()
	collector.ImportTemplate()
	collector.ImportExtensions()
	collector.WebhookLoad()
}
