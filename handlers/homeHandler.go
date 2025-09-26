package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
	<head>
		<title>Mini Blockchain</title>
	</head>
	<body>
		<h1>🚀 Bienvenue sur ma Mini Blockchain</h1>
		<p>Choisis une action :</p>
		<ul>
			<li><a href="/transactions">➕ Ajouter une transaction</a></li>
			<li>
				<form action="/mine" method="POST" style="display:inline;">
					<button type="submit">⚒️ Valider un bloc</button>
				</form>
			</li>
			<li><a href="/blocks">📜 Voir la blockchain</a></li>
		</ul>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}
