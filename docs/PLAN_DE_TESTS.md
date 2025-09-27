# Plan de tests – Projet Blockchain Go

## 1. Objectif et périmètre
Vérifier le bon fonctionnement de l’application **Blockchain** déployée avec Docker et accessible via un reverse proxy.  
Les tests portent sur :
- Les endpoints `/health`, `/transactions`, `/mine`, `/blocks`
- La validation des entrées (transactions valides / invalides)
- La gestion des erreurs (payload JSON incorrect)
- L’intégration du pipeline CI/CD (exécution automatisée des tests)

---

## 2. Environnement
- **Langage** : Go 1.25
- **Frameworks** : `net/http`, `testing`
- **Conteneurs** : Docker + Docker Compose
- **CI/CD** : GitHub Actions
- **Endpoints testés** : `http://localhost:8080`

---

## 3. Cas de tests

| ID   | Description                                 | Pré-requis             | Étapes                                                                 | Résultat attendu                  |
|------|---------------------------------------------|------------------------|------------------------------------------------------------------------|-----------------------------------|
| T1   | Vérifier endpoint santé `/health`           | Serveur démarré        | `GET /health`                                                          | Retourne `200 OK` + `OK`          |
| T2   | Récupération liste transactions (vide)      | Serveur démarré        | `GET /transactions` (JSON)                                             | Retourne `200 OK` + `[]`          |
| T3   | Ajout transaction invalide                  | Serveur démarré        | `POST /transactions` `{}`                                              | Retourne `400 Bad Request`        |
| T4   | Ajout transaction valide                    | Serveur démarré        | `POST /transactions` JSON `{sender, receiver, amount>0}`               | Retourne `201 Created` + transaction |
| T5   | Minage sans transaction                     | Liste vide             | `POST /mine`                                                           | Retourne `200 OK` + “Aucune transaction à valider” |
| T6   | Minage avec transaction                     | Une transaction ajoutée| `POST /mine`                                                           | Retourne `200 OK` + bloc créé     |
| T7   | Récupération de la blockchain               | Après minage           | `GET /blocks`                                                          | Retourne chaîne non vide           |
| T8   | Sécurité : payload JSON malformé            | Serveur démarré        | `POST /transactions` avec JSON invalide                                | Retourne `400 Bad Request`        |

---

## 4. Résultats
Extrait GitHub Actions :  
```bash
=== RUN   TestHealthHandler
--- PASS: TestHealthHandler
=== RUN   TestTransactionsHandler
--- PASS: TestTransactionsHandler
=== RUN   TestMineHandler
--- PASS: TestMineHandler
PASS
coverage: 64.7% of statements
