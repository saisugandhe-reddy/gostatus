# gostatus

`gostatus` is a small Go CLI that checks URLs concurrently and prints their HTTP status, latency, and any request errors.

It is simple enough to understand quickly, but structured enough to make a decent starter repository for a GitHub profile.

## Features

- Concurrent URL checks
- Configurable timeout
- Clean table output
- Unit tests for core behavior

## Example

```bash
go run ./cmd/gostatus --timeout 3s https://golang.org https://github.com https://example.com
```

Example output:

```text
URL                    STATUS   LATENCY   ERROR
https://golang.org     200      142ms
https://github.com     200      181ms
https://example.com    200      96ms
```

## Project layout

```text
cmd/gostatus           CLI entrypoint
internal/checker       URL checking logic
```

## Run locally

```bash
go run ./cmd/gostatus https://example.com
```

## Test

```bash
go test ./...
```

## Push to GitHub

Create a new empty GitHub repository named `gostatus`, then run:

```bash
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/saisugandhe-reddy/gostatus.git
git push -u origin main
```
