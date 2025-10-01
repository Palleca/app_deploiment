package handlers

import (
	"fmt"
	"net/http"
)

// GET / → page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
	<head>
		<title>Mini Blockchain</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; background: #f9f9f9; }
			.container { background: #fff; padding: 30px; border-radius: 8px; 
			             box-shadow: 0 2px 6px rgba(0,0,0,0.1); max-width: 600px; margin: auto; }
			h1 { color: #007BFF; text-align: center; }
			ul { list-style-type: none; padding: 0; }
			li { margin: 15px 0; }
			a, button { 
				text-decoration: none; 
				color: white; 
				background: #007BFF; 
				border: none; 
				padding: 12px 20px; 
				border-radius: 4px; 
				cursor: pointer;
				font-size: 16px;
			}
			a:hover, button:hover { background: #0056b3; }
			form { display: inline; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>🚀 Bienvenue sur ma Mini Blockchain</h1>
			<p style="text-align:center;">Choisis une action :</p>
			<ul>
				<li><a href="/transactions">➕ Ajouter une transaction</a></li>
				<li>
					<form action="/mine" method="POST">
						<button type="submit">⚒️ Valider un bloc</button>
					</form>
				</li>
				<li><a href="/blocks">📜 Voir la blockchain</a></li>
			</ul>
		</div>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}
