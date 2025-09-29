package main

import (
	"app_deploiment/pkg"
	"log"
	"net/http"
	"os"
)

func main() {
	r := pkg.Router()

	// Récupère le port depuis l'environnement, sinon par défaut 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Serveur démarré sur http://localhost:%s", port)

	// Bloque tant que le serveur tourne, et log l'erreur si jamais il plante
	log.Fatal(http.ListenAndServe(":"+port, r))
}
