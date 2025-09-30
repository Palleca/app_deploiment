# 1. Vérifier que le code Go tourne en local
Write-Host "==> 1. Lancement direct de l'app Go"
go run main.go

# (Ouvrir un autre terminal pour tester avec curl)
# curl.exe http://localhost:8080/blocks

# 2. Lancer les tests unitaires Go
Write-Host "==> 2. Tests unitaires"
go test ./...

# 3. Construire l'image Docker de l'API blockchain
Write-Host "==> 3. Dockerisation de l'API blockchain"
docker build -t app_blockchain:latest .

# 4. Vérifier que l'image tourne en local
Write-Host "==> 4. Vérification en local"
docker run -d -p 8080:8080 --name blockchain_test app_blockchain:latest
curl.exe http://localhost:8080/blocks
docker stop blockchain_test
docker rm blockchain_test

# 5. Construire l'image du proxy
Write-Host "==> 5. Dockerisation du proxy"
docker build -t app_proxy:latest ./proxy

# 6. Lancer tout via Docker Compose (proxy + API)
Write-Host "==> 6. Orchestration avec Docker Compose"
docker-compose up -d --build
curl.exe http://localhost:8080/blocks

# 7. Pousser l'image sur Docker Hub
Write-Host "==> 7. Partage via Docker Hub"
docker tag app_blockchain:latest tonDockerHub/app_blockchain:latest
docker tag app_proxy:latest tonDockerHub/app_proxy:latest
docker push tonDockerHub/app_blockchain:latest
docker push tonDockerHub/app_proxy:latest

# 8. (Automatisé via CI/CD)
Write-Host "==> 8. CI/CD (GitHub Actions fait build/test/push automatiquement)"

# 9. Déploiement sur un serveur (simulation locale)
Write-Host "==> 9. Déploiement depuis Docker Hub"
docker-compose down
docker pull tonDockerHub/app_blockchain:latest
docker pull tonDockerHub/app_proxy:latest
docker-compose up -d
curl.exe http://localhost:8080/blocks
