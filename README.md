# Renco Boilerplate

## Getting started

This is built on top of [Go Echo](https://echo.labstack.com) Golang Framework.

## Dependencies

There is some dependencies that we used in this skeleton:

- [Go Echo](https://echo.labstack.com/) [Go Framework]
- [Viper](https://github.com/spf13/viper) [Go Configuration]
- [Cobra](https://github.com/spf13/cobra) [Go Modern CLI]
- [Logrus Logger](https://github.com/sirupsen/logrus) [Go Logger]
- [Goose Migration](https://github.com/pressly/goose) [Go Migration]
- [OpenTelemetry](https://pkg.go.dev/go.opentelemetry.io/otel) [OpenTelemetry Tracer] - Upcoming

## Requirement

- Golang version 1.23 or latest
- Database PostgreSQL
- Firebase Auth

## Usage

### Installation

install required dependencies

```bash
make install
```

### Run Service

run current service after all dependencies installed

```bash
make start-http
```

### Build Service

run this command to build service to binary

```bash
make build
```

run to start service from binary of application

```bash
make run-http
```

# Database Migration Management

This document provides instructions for managing database migrations using the provided `Makefile`. The `Makefile` includes commands for various migration operations, including applying migrations, rolling them back, and creating new migrations.

## Migration Commands

### Show Help

```bash
make help
```

### Migrate Up

Apply all pending migrations to the database.

```bash
make migrate-up
```

### Create Migration

```bash
make migrate-create NAME=<migration_name> [TYPE=sql]

# example : make migrate-create NAME=users TYPE=sql
```

## Health Check Endpoint

### Go to Endpoint

```bash
{base_url}/up
```

Expected Response:

```json
{
  "ref_id": "z0LMVDmRkV",
  "code": 200,
  "status": "OK",
  "data": {
    "message": "Waras!"
  },
  "timestamp": "2024-09-19T15:31:20.009226+07:00"
}
```
