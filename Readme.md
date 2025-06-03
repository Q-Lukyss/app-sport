# App Sport

Suivi du Poids et du Sport   

Go Gin + Next + PostgreSQL   


## Recup les dépendances

### Go

dans /api
```
go mod tidy
```

### Node

dans /frontend
```
npm install
```

## Setup Base de donnée

A la racine 
```
docker compose up -d --build
```

Supprimer le volume 
```
docker compose down -v
```

## Lancer Gin

dans /api
```
go run main.go
```

## Lancer Next

dans /frontend

```
npm run dev
```

## Swagger et Routes API

pour mettre a jour la doc swagger 

dans /api
```
swagg init
```

la doc est dispo a /swagger/index.html

## 🔄 Hot reload avec Air (optionnel)

- Installer Air :  
  `go install github.com/cosmtrek/air@latest`

- Lancer en mode dev :  
  `air` (depuis le dossier `api/`)

- Config facultative : voir `.air.toml`

```
# .air.toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_dir = ["assets", "tmp", "vendor"]

[log]
  time = true
```