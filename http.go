package main

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"os"
	"time"
)

func serveSite() {

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(flgHost),
		Cache:      autocert.DirCache("certs"),
	}

	router := mux.NewRouter()
	router.HandleFunc("/", handleRoot).Methods("GET")
	router.PathPrefix("/x/").HandlerFunc(handlePost).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(flgStatic))))

	var srv *http.Server

	if flgSecure {
		srv = &http.Server{
			Handler: router,
			Addr:    ":https",

			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}
	} else {
		srv = &http.Server{
			Handler: router,
			Addr:    ":8080",

			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
	}

	if flgSecure {
		go http.ListenAndServe(":http", certManager.HTTPHandler(nil))
		srv.ListenAndServeTLS("", "")
	} else {
		err := srv.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/x/", 302)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World 3"))
}
