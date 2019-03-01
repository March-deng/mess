package web

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"testing"

	"golang.org/x/crypto/acme/autocert"
)

func TestTLSServe(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS user! Your config: %+v", r.TLS)
	})
	log.Fatal(http.Serve(autocert.NewListener("example.com"), mux))
}

func TestCPUNum(t *testing.T) {
	n := runtime.NumCPU()
	log.Println(n)
}
