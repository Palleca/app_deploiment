package main

import (
	"app_deploiment/pkg"
	"log"
	"net/http"
)

func main() {
	r := pkg.Router()
	log.Println("🚀 Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
