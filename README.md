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
go run . observations add "Observation fictive"
```

Les observations sont actuellement stockées en mémoire. Elles ne sont pas encore conservées après la fin du programme.

## Vérification

```sh
go fmt ./...
go test ./...
go vet ./...
```