# Renco Boilerplate

## Getting started

This is built on top of [Go Echo](https://echo.labstack.com) Golang Framework.

## Dependencies

There is some dependencies that we used in this skeleton:

- [Go Echo](https://echo.labstack.com/) [Go Framework]
- [Viper](https://github.com/spf13/viper) [Go Configuration]
- [Cobra](https://github.com/spf13/cobra) [Go Modern CLI]
- [Logrus Logger](https://github.com/sirupsen/logrus) [Go Logger]
- [Goose Migration](https://github.com/pressly/goose) [Go Migration] - Upcoming
- [Gobreaker](https://github.com/sony/gobreaker) [Go Circuit Breaker] - Upcoming
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
