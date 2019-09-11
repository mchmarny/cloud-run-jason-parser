package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	ev "github.com/mchmarny/gcputil/env"
)

var (
	logger = log.New(os.Stdout, "", 0)
	port   = ev.MustGetEnvVar("PORT", "8080")
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/select", selectHandler)
	mux.HandleFunc("/find", findHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

	hostPort := net.JoinHostPort("0.0.0.0", port)
	server := &http.Server{Addr: hostPort, Handler: mux}
	logger.Printf("Server starting: %s \n", hostPort)
	logger.Fatal(server.ListenAndServe())
}
