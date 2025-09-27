
---

## 📄 `docs/DEPLOIEMENT.md`
```markdown
# Procédure de déploiement – Projet Blockchain Go

## 1. Prérequis
- Docker + Docker Compose installés
- Accès au registre Docker Hub : `palleca/app_deploiment` et `palleca/app_deploiment-proxy`
- Variables d’environnement (secrets) :
  - `DOCKER_USERNAME`
  - `DOCKER_PASSWORD`

---

## 2. Build & push (CI/CD)
La CI GitHub Actions :
1. Exécute les tests Go (`go test ./...`)
2. Construit les images Docker
3. Push sur Docker Hub :
   - `palleca/app_deploiment:latest`
   - `palleca/app_deploiment-proxy:latest`

---

## 3. Déploiement manuel (local ou serveur)
```bash
# Récupération des dernières images
docker pull palleca/app_deploiment:latest
docker pull palleca/app_deploiment-proxy:latest

# Lancement via Docker Compose
docker compose up -d
