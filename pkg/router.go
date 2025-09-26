package pkg

import (
	"app_deploiment/handlers"

	"github.com/gorilla/mux"
)

// Router expose les routes de l’application pour être utilisées dans main.go et dans les tests
func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/blocks", handlers.GetBlocksHandler).Methods("GET")
	r.HandleFunc("/transactions", handlers.GetTransactionsHandler).Methods("GET")
	r.HandleFunc("/transactions", handlers.CreateTransactionHandler).Methods("POST")
	r.HandleFunc("/mine", handlers.MineBlockHandler).Methods("POST")
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/reset", handlers.ResetHandler).Methods("POST", "GET")

	return r
}
