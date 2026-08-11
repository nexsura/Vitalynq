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
go run . privacy
go run . observations list
go run . obs list
go run . observations add "Observation fictive"
go run . obs add "Observation fictive"
go run . --db test.db observations list
go run . db path
go run . --db test.db db path
go run . observations add --date 2026-07-29 "Observation fictive"
go run . profile set "Profil fictif"
go run . profile show
go run . measurements list
go run . measurements add poids 72.5 kg "test fictif" "saisie manuelle" "saisie manuelle"
go run . measurements add --date 2026-07-29 poids 72.5 kg "test fictif" "saisie manuelle" "saisie manuelle"
go run . appointments list
go run . appointments add 2026-07-29 "consultation fictive" rendez-vous "cabinet fictif" "saisie manuelle"
go run . summary
go run . export
```

Les observations sont actuellement stockées dans un fichier SQLite local `vitalynq.db`.
L'option `--db` permet de choisir un autre fichier SQLite local.

La commande `export` affiche les données locales au format JSON dans le terminal. Elle n'envoie aucune donnée vers un serveur, un cloud ou un service externe.

La commande `privacy` affiche les garanties et limites de confidentialité de Vitalynq : stockage local, absence de cloud, absence de télémétrie et responsabilité de protéger les fichiers locaux.

## Vérification

```sh
go fmt ./...
go test ./...
go vet ./...
```