package handlers

import (
	"app_deploiment/core"
	"encoding/json"
	"fmt"
	"net/http"
)

// POST /mine → valide les transactions et ajoute un bloc
func MineBlockHandler(w http.ResponseWriter, r *http.Request) {
	if len(core.Transactions) == 0 {
		http.Error(w, "Aucune transaction à valider", http.StatusBadRequest)
		return
	}

	newBlock := core.MineBlock()

	// JSON si demandé
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(newBlock)
		return
	}

	// Sinon HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := fmt.Sprintf(`
	<html>
	<head>
		<title>Nouveau Bloc</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; background: #f9f9f9; }
			.container { background: #fff; padding: 20px; border-radius: 8px; 
			             box-shadow: 0 2px 6px rgba(0,0,0,0.1); max-width: 600px; }
			h1 { color: #28a745; }
			pre { background: #eee; padding: 10px; border-radius: 4px; }
			a { display: inline-block; margin-top: 20px; text-decoration: none; color: #007BFF; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>✅ Nouveau bloc validé</h1>
			<p><b>Index :</b> %d</p>
			<p><b>Hash :</b> %s</p>
			<p><b>Précédent :</b> %s</p>
			<h2>Transactions validées :</h2>
			<pre>%+v</pre>
			<a href="/">⬅️ Retour à l'accueil</a>
		</div>
	</body>
	</html>
	`, newBlock.Index, newBlock.Hash, newBlock.PreviousHash, newBlock.Transactions)

	fmt.Fprint(w, html)
}
