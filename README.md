# ToDo App Backend Maxi

TODO APP BACKEND - MAXI

## Prelude

This is a backend service for ToDo App

## Requirements

- Golang 1.16 or higher

## Note:

> Currently  only tested on MacOS, if you are using any OS beside MacOS, feel free to contribute to this README

## Setup

### Install Golang dependencies

On console, run following command to install required golang dependencies

```shell
go mod tidy
```

Rename config.yaml.example -> config.yaml (modify accordingly)
```shell
cp cp .env.example .env
```

## Run

### Start HTTP Server
```shell
go run main.go
```