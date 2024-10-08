# Amartha Test Repository

This repository is structured to enable rapid development, with a primary focus on implementing the solution logic for the state engine of a P2P lending platform.

## Setup

To set up the environment:

- Install goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

- Copy .env.example to .env then fill in the .env file

## Migrations

```bash
goose -dir migration/ mysql "user:password@tcp(localhost:3306)/test_amartha?parseTime=true" up
```
