package handlers

import (
	"app_deploiment/core"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GET /transactions → formulaire HTML OU JSON
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	// Mode JSON (tests, API)
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(core.Transactions)
		return
	}

	// Sinon mode HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
	<html>
	<head>
		<title>Transactions</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; background: #f9f9f9; }
			h1 { color: #333; }
			form { background: #fff; padding: 20px; border-radius: 8px;
			       box-shadow: 0 2px 6px rgba(0,0,0,0.1); width: 300px; margin-bottom: 30px; }
			label { font-weight: bold; display: block; margin-top: 10px; }
			input { width: 100%; padding: 8px; margin-top: 5px; border: 1px solid #ccc; border-radius: 4px; }
			button { margin-top: 15px; width: 100%; padding: 10px; background: #007BFF; color: white;
			         border: none; border-radius: 4px; cursor: pointer; }
			button:hover { background: #0056b3; }
			table { border-collapse: collapse; width: 60%; margin-top: 20px; background: #fff; }
			th, td { border: 1px solid #ddd; padding: 10px; text-align: center; }
			th { background: #007BFF; color: white; }
			tr:nth-child(even) { background: #f2f2f2; }
			a { text-decoration: none; color: #007BFF; }
		</style>
	</head>
	<body>
		<h1>Ajouter une Transaction</h1>
		<form action="/transactions" method="POST">
			<label>Expéditeur :</label>
			<input type="text" name="sender" required>

			<label>Destinataire :</label>
			<input type="text" name="receiver" required>

			<label>Montant :</label>
			<input type="number" step="0.01" name="amount" required>

			<button type="submit">Ajouter la transaction</button>
		</form>

		<h2>Transactions en attente</h2>
		<table>
			<tr><th>Expéditeur</th><th>Destinataire</th><th>Montant</th></tr>
	`

	if len(core.Transactions) == 0 {
		html += `<tr><td colspan="3">Aucune transaction en attente</td></tr>`
	} else {
		for _, tx := range core.Transactions {
			html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%.2f</td></tr>",
				tx.Sender, tx.Receiver, tx.Amount)
		}
	}

	html += `
		</table>
		<br>
		<a href="/">⬅️ Retour à l'accueil</a>
	</body>
	</html>
	`

	fmt.Fprint(w, html)
}

// POST /transactions → JSON ou Formulaire
func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Cas JSON (tests, API)
	if r.Header.Get("Content-Type") == "application/json" {
		var tx core.Transaction
		if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
			http.Error(w, "Payload JSON invalide", http.StatusBadRequest)
			return
		}
		if tx.Sender == "" || tx.Receiver == "" || tx.Amount <= 0 {
			http.Error(w, "Transaction invalide (sender/receiver/amount)", http.StatusBadRequest)
			return
		}

		core.AddTransaction(tx)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status": "Transaction ajoutée avec succès",
			"tx":     tx,
		})
		return
	}

	// Cas Formulaire HTML
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur parsing formulaire", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil || amount <= 0 {
		http.Error(w, "Montant invalide", http.StatusBadRequest)
		return
	}

	tx := core.Transaction{
		Sender:   r.FormValue("sender"),
		Receiver: r.FormValue("receiver"),
		Amount:   amount,
	}

	core.AddTransaction(tx)
	http.Redirect(w, r, "/transactions", http.StatusSeeOther)
}
