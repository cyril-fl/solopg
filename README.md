# SoloPG

CLI Go pour piloter une session de solo RPG dans le terminal.

## Versioning

Le projet utilise des metadonnees de build centralisees dans `internal/buildinfo`.

- Sans `ldflags`, la commande `version` utilise des valeurs par defaut et tente un fallback via `runtime/debug`.
- Avec `ldflags`, les valeurs de version, commit et date sont injectees au build.

### Commandes recommandees

```bash
make run
make build
./solopg version
./solopg --version
```

### Build manuel

```bash
go build -ldflags "-X 'solopg/internal/buildinfo.AppName=solopg' -X 'solopg/internal/buildinfo.Version=v0.0.3' -X 'solopg/internal/buildinfo.Commit=$(git rev-parse --short HEAD)' -X 'solopg/internal/buildinfo.Date=$(date -u +%Y-%m-%dT%H:%M:%SZ)'" .
```
