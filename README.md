# API - Magic Number

## Principe
Au lancement du serveur, un nombre magique (aléatoire, entre 1 et 100) est généré. Vous devez le deviner.

## Appeler l'API
S'inscrire au jeu.
```curl
curl -X POST http://localhost:8080/join -d '{"name": "toto"}' -H "Content-Type: application/json"
```

Se désinscrire au jeu.
```curl
curl -X POST http://localhost:8080/leave -d '{"name": "toto"}' -H "Content-Type: application/json"
```

Lancer le jeu/Générer un nouveau nombre magique.
```curl
curl -X POST http://localhost:8080/start -H "Content-Type: application/json"
```

Tenter un guess.
```curl
curl -X POST http://localhost:8080/play -d '{"name": "toto", "guess": 25}' -H "Content-Type: application/json"
```
