# go_etablissement_ms

POC d'un serveur avec API REST des établissements en golang, se base sur la lecture de data.txt.
Utilise go-redis, gin.

## Prérequis

- golang
- podman

## Utilisation


Générer un data.txt
```
cd writer
go build
./main -size=100
```

Lancer le serveur
```
sudo podman run -d --name redis_server -p 6379:6379 docker.io/redis
go run .
```

Browser

http://localhost:8080/etablissement/:id

Example: 

http://localhost:8080/etablissement/0000010

## Commandes utiles

```
// liste des conteneurs
podman ps
// arrêter un conteneur
podman stop <conteneurID>
```