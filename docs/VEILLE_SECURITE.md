
---

## 📄 `SECURITY.md`
```markdown
# Sécurité – Projet Blockchain Go

## 1. Risques identifiés
- Injection de données malformées
- Transactions invalides (montant ≤ 0, champs vides)
- Attaques réseau (DoS, brute-force sur le proxy)
- Exposition des secrets (tokens, identifiants Docker)

---

## 2. Mesures mises en place
- **Validation stricte** des champs transaction côté API
- **CI/CD sécurisé** : secrets gérés via GitHub Secrets (non commités)
- **Conteneurs isolés** via Docker Compose (apps + proxy)
- **Reverse proxy** : prévu pour gérer TLS (certificats dans `reverse_proxy/certs/`)

---

## 3. Améliorations recommandées
- Activer TLS (terminaison HTTPS) au niveau du proxy
- Ajouter des **HTTP security headers** (CSP, HSTS, etc.)
- Mettre en place une **limite de requêtes** (rate limiting) sur le proxy
- Ajouter un **scan de vulnérabilités** automatisé :
  - `go vet`, `govulncheck`
  - `trivy image palleca/app_deploiment:latest`

---

## 4. Veille
- Surveiller les alertes de sécurité des dépendances Go (`go list -m -u all`)
- Vérifier régulièrement les CVE sur Debian slim et Golang
- Consulter OWASP Top 10 (notamment API Security)
