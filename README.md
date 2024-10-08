# TEST AMARTHA

This repository is designed with a simple structure to facilitate fast development. The primary focus is on the implementation of solution logic for state engine for P2P lending.

## Setup

- Install goose -> go install github.com/pressly/goose/v3/cmd/goose@latest
- Copy .env.example to .env
- Fill in the .env file

## Migrations

- goose -dir migration/ mysql "root:password@tcp(localhost:3306)/test_amartha?parseTime=true" up
