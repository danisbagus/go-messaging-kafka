# go-messaging-kafka
Simple golang messaging app with Kafka

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Docker-compose](https://docs.docker.com/compose/) for running Kafka.

## Setup

Create Kafka container

```bash
docker-compose up
```

## Run the service

Prepare necessary environemt by rename .env.example to .env

```bash
CONFIG_SMTP_HOST="smtp.gmail.com"
CONFIG_SMTP_PORT=587
CONFIG_SENDER_NAME=
CONFIG_AUTH_EMAIL=
CONFIG_AUTH_PASSWORD=
KAFKA_BROKER_1=
TOPIC_TRANSACTION=
GROUP_ID_TRANSACTION=
```

Get go packages

```bash
go get .
```

Run transaction service

```bash
go run transaction-service/main.go
```

Run mail service

```bash
go run mail-service/main.go
```
