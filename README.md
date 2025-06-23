# Sedis

**Statut du projet : terminé**<br>

Sedis est une implémentation minimaliste de Redis, conçue pour gérer des opérations clé-valeur basiques, avec persistance sur fichier et une interface en ligne de commande.<br>

## Architecture

- **CLI intégré** (port 8085) : interface locale pour gérer et tester les commandes manuellement
- **Serveur TCP** (port 8086) : permet la connexion de clients externes (ex. : backends applicatifs)

## Compilation

```
docker-compose build
docker-compose up -d
```

## Utilisation
```
docker-compose run --rm sedis_manager ./sedis_manager [COMMAND]
```

**Commandes:**<br>
```
TTL     key                           # Affiche le temps restant avant expiration
EXISTS  key                           # Vérifie si une clé existe
LIST                                  # Liste toutes les clés
SET     key value [NX|XX] [EX sec]    # Définit une valeur avec options
GET     key                           # Récupère la valeur d'une clé
DEL     key                           # Supprime une clé
FLUSHALL                              # Supprime toutes les clés
SAVE    [fileName]                    # Sauvegarde dans un fichier
LOAD    [fileName]                    # Charge depuis un fichier
```

## Fonctionnalités
- Stockage clé-valeur en mémoire
- Gestion d'expiration (`TTL`, `EX`)
- Conditions d’écriture (`NX`, `XX`)
- Persistance via `SAVE` / `LOAD`
- Connexion TCP pour clients distants
- CLI locale pour l'administration et les tests