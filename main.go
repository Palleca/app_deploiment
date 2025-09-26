package main

import (
	"app_deploiment/pkg"
	"log"
	"net/http"
)

func main() {
	r := pkg.Router()
	log.Println("🚀 Serveur démarré sur http://localhost:8080")

	// Bloque tant que le serveur tourne, et log l'erreur si jamais il plante
	log.Fatal(http.ListenAndServe(":8080", r))
}
