# coverage.ps1 - Génération et affichage de la couverture de tests Go

# Définir la racine du projet (dossier courant)
$projectRoot = Get-Location

# Définir le dossier de coverage
$coverageDir = Join-Path $projectRoot "tests\coverage"

# Créer le dossier s'il n'existe pas
if (-Not (Test-Path $coverageDir)) {
    New-Item -ItemType Directory -Path $coverageDir | Out-Null
}

# Définir les fichiers coverage
$coverageOut = Join-Path $coverageDir "coverage.out"
$coverageHtml = Join-Path $coverageDir "coverage.html"

Write-Host "➡️  Lancement des tests avec couverture (excluant main.go et logger.go)..."

# On limite le coverage aux packages core, handlers et pkg
go test ./... -coverpkg=./core,./handlers,./pkg -coverprofile="$coverageOut"

Write-Host "➡️  Génération du rapport HTML..."
go tool cover -html="$coverageOut" -o "$coverageHtml"

Write-Host "➡️  Résumé de la couverture :"
go tool cover -func="$coverageOut"

Write-Host "➡️  Ouverture du rapport HTML..."
Start-Process "$coverageHtml"
