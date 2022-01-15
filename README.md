# Marketplace Engine
REST APIs for Web3 Auth and Smart Contract Functionalities

[![.github/workflows/test.yml](https://github.com/TheLazarusNetwork/marketplace-engine/actions/workflows/test.yml/badge.svg)](https://github.com/TheLazarusNetwork/marketplace-engine/actions/workflows/test.yml)

[![Lint](https://github.com/TheLazarusNetwork/marketplace-engine/actions/workflows/lint.yml/badge.svg)](https://github.com/TheLazarusNetwork/marketplace-engine/actions/workflows/lint.yml)

# Getting Started

## Postgres for development
```bash
docker run --name="marketplace" --rm -d -p 5432:5432 \
-e POSTGRES_PASSWORD=revotic \
-e POSTGRES_USER=revotic \
-e POSTGRES_DB=marketplace \
postgres -c log_statement=all
```

## Steps to get started
- Run `go get ./...` to install dependencies
- Set up env variables or create `.env` file as per [`.env-sample`](https://github.com/TheLazarusNetwork/marketplace-engine/blob/main/.env-sample) file
- Run `go test ./...` to make sure setup is working
- Run `go run main.go` to start server
