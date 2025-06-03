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