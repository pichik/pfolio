package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pichik/pfolio/src/auth"
	"github.com/pichik/pfolio/src/data"
	"github.com/pichik/pfolio/src/database"
	"github.com/pichik/pfolio/src/misc"
	request "github.com/pichik/pfolio/src/requests"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"golang.org/x/crypto/acme/autocert"
)

func setHeaders(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if !auth.IsAuthed(r) {
		var headers string
		for key, values := range r.Header {
			for _, value := range values {
				headers += fmt.Sprintf("%s: %s\n", key, value)
			}
		}
		misc.RequestLog.Printf("%-17s %9s %s\n%s", "["+strings.Split(r.RemoteAddr, ":")[0]+"]", "["+r.Method+"]", r.RequestURI, headers)
	}

	origin := r.Header.Get("Origin")

	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		return
	}
	next(w, r)
}

func checkAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// if !auth.IsAuthed(r) {
	// 	return
	// }
	next(w, r)
}
func login(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func main() {
	loadFiles()
	database.Opendb(data.DataDir)
	go UpdateStocks()

	r := mux.NewRouter()
	n := negroni.Classic()

	n.Use(negroni.NewStatic(http.Dir(data.AssetsDir)))
	n.UseFunc(setHeaders)

	setupDataRoutes(r)

	r.PathPrefix("/login").Handler(negroni.New(
		negroni.HandlerFunc(auth.Authenticate),
		negroni.HandlerFunc(checkAuth),
	))

	n.UseHandler(r)

	server, certManager := setupServer(n)
	go http.ListenAndServe(":8000", certManager.HTTPHandler(n))
	log.Fatal(server.ListenAndServeTLS("", ""))
}

func setupDataRoutes(r *mux.Router) {
	dataRoutes := mux.NewRouter().PathPrefix("/").Subrouter()
	dataRoutes.HandleFunc("/import-xtb", request.ImportXTB).Methods("POST")
	dataRoutes.HandleFunc("/add-to-wlist", request.ImportWlist).Methods("POST")
	dataRoutes.HandleFunc("/stock-update", request.StocksUpdate).Methods("GET")
	dataRoutes.HandleFunc("/login", auth.Login).Methods("GET")

	r.PathPrefix("/").Handler(negroni.New(
		negroni.HandlerFunc(checkAuth),
		negroni.Wrap(dataRoutes),
	))
}

func setupServer(n *negroni.Negroni) (*http.Server, autocert.Manager) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("pfolio." + data.Config.Hostname),
		Cache:      autocert.DirCache(data.CertDir),
	}

	server := &http.Server{
		Addr:      ":8443",
		Handler:   n,
		TLSConfig: certManager.TLSConfig(),
	}
	return server, certManager
}

func loadFiles() {
	misc.ImportLogs()
	//datacenter.CreateDatabase()
}

func UpdateStocks() {
	ticker30Sec := time.NewTicker(3 * time.Second)
	ticker24Hours := time.NewTicker(24 * time.Hour)
	request.UpdateEchangeRates()
	for {
		select {
		case <-ticker30Sec.C:
			database.ReadDatabase()

		case <-ticker24Hours.C:
			request.UpdateEchangeRates()
		}
	}
}
