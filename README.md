# Vitalynq

Vitalynq is a local CLI application under construction for organizing personal health data.

It organizes data. It does not provide diagnosis, recommend treatments, or replace a healthcare professional.

## Running

```sh
go run .
```

## Available Commands

```sh
go run . help
go run . version
go run . about
go run . privacy
go run . limitations
go run . observations list
go run . obs list
go run . observations add "Fictive observation"
go run . obs add "Fictive observation"
go run . --db test.db observations list
go run . db path
go run . --db test.db db path
go run . db info
go run . --db test.db db info
go run . db check
go run . --db test.db db check
go run . observations add --date 2026-07-29 "Fictive observation"
go run . profile set "Fictive profile"
go run . profile show
go run . measurements list
go run . measurements add weight 72.5 kg "fictive test" "manual entry" "manual entry"
go run . measurements add --date 2026-07-29 weight 72.5 kg "fictive test" "manual entry" "manual entry"
go run . appointments list
go run . appointments add 2026-07-29 "fictive consultation" appointment "fictive office" "manual entry"
go run . summary
go run . export
```

Observations are currently stored in a local SQLite file named `vitalynq.db`.
The `--db` option can select another local SQLite file.

The `export` command prints local data as JSON in the terminal. It does not send data to a server, cloud service, or external service.

The `privacy` command shows Vitalynq privacy guarantees and limits: local storage, no cloud, no telemetry, and the user's responsibility to protect local files.

The `limitations` command shows Vitalynq functional limits: data organization only, no diagnosis, no treatment recommendation, no prediction, and no replacement for a healthcare professional.

## Verification

```sh
go fmt ./...
go test ./...
go vet ./...
```
