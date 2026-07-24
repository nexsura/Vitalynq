# Vitalynq

Vitalynq est une application CLI locale en cours de construction pour organiser des données personnelles de santé.

Elle organise des données. Elle ne pose pas de diagnostic, ne recommande pas de traitement et ne remplace pas un professionnel de santé.

## Exécution

```sh
go run .
```

## Commandes disponibles

```sh
go run . help
go run . version
go run . about
go run . observations list
go run . obs list
go run . observations add "Observation fictive"
go run . obs add "Observation fictive"
go run . --db test.db observations list
```

Les observations sont actuellement stockées dans un fichier SQLite local `vitalynq.db`.
L'option `--db` permet de choisir un autre fichier SQLite local.

## Vérification

```sh
go fmt ./...
go test ./...
go vet ./...
```