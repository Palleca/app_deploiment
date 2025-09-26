package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
)

var counter uint32

func main() {
	// Récupère la liste des backends (séparés par des virgules)
	backendEnv := os.Getenv("BACKENDS")
	if backendEnv == "" {
		log.Fatal("❌ Aucune URL backend fournie dans BACKENDS")
	}

	backendURLs := strings.Split(backendEnv, ",")
	var proxies []*httputil.ReverseProxy

	for _, b := range backendURLs {
		target, err := url.Parse(strings.TrimSpace(b))
		if err != nil {
			log.Fatalf("❌ Erreur parsing backend URL %s : %v", b, err)
		}
		proxies = append(proxies, httputil.NewSingleHostReverseProxy(target))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// round-robin
		idx := int(atomic.AddUint32(&counter, 1)) % len(proxies)
		log.Printf("➡️ %s %s → backend %d", r.Method, r.URL.Path, idx+1)
		proxies[idx].ServeHTTP(w, r)
	})

	log.Println("🚀 Proxy démarré sur :80 avec backends :", backendURLs)
	log.Fatal(http.ListenAndServe(":80", nil))
}
