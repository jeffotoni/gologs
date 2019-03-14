#!/bin/bash
GOOS=linux go build -ldflags="-s -w" -o gologs main.go

### GMAIL SMTP
export GMAIL_PASSWORD=xxxxxx
export GMAIL_USER=xxxxxx
export EMAIL_NOTIFIY=xxxxxx

#### DB POSTGRES
export DB_NAME=xxxxxx
export DB_HOST=localhost
export DB_USER=xxxxxx
export DB_PASSWORD=xxxxxx
export DEBUG=false
export DEBUG_REQ=10000
./gologs
