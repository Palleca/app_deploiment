package handlers

import (
	"app_deploiment/core"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GET /blocks → affiche la blockchain
func GetBlocksHandler(w http.ResponseWriter, r *http.Request) {
	// JSON si demandé
	if strings.Contains(r.Header.Get("Accept"), "application/json") {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(core.Blockchain)
		return
	}

	// Sinon HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
	<html>
	<head>
		<title>Blockchain</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; background: #f9f9f9; }
			h1 { color: #333; }
			.block { background: #fff; padding: 20px; margin-bottom: 20px; border-radius: 8px;
			         box-shadow: 0 2px 6px rgba(0,0,0,0.1); }
			.tx { margin-left: 20px; }
			a, form button { text-decoration: none; color: white; background: #007BFF; 
			                 border: none; padding: 10px 15px; border-radius: 4px; cursor: pointer; }
			form { display: inline; }
		</style>
	</head>
	<body>
		<h1>📜 Blockchain</h1>
	`

	for _, block := range core.Blockchain {
		html += fmt.Sprintf(`
		<div class="block">
			<h2>Bloc %d</h2>
			<p><b>Hash :</b> %s</p>
			<p><b>Précédent :</b> %s</p>
			<h3>Transactions :</h3>
			`, block.Index, block.Hash, block.PreviousHash)

		if len(block.Transactions) == 0 {
			html += "<p>Aucune transaction dans ce bloc.</p>"
		} else {
			for _, tx := range block.Transactions {
				html += fmt.Sprintf("<div class='tx'>%s → %s : %.2f</div>", tx.Sender, tx.Receiver, tx.Amount)
			}
		}

		html += "</div>"
	}

	html += `
		<form action="/reset" method="POST">
			<button type="submit">🔄 Réinitialiser la Blockchain</button>
		</form>
		<br><br>
		<a href="/">⬅️ Retour à l'accueil</a>
	</body>
	</html>
	`

	fmt.Fprint(w, html)
}
