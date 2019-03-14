#!/bin/bash
GOOS=linux go build -ldflags="-s -w" -o gologs main.go

### GMAIL SMTP
export GMAIL_PASSWORD=xxxxxx
export GMAIL_USER=xxxxxx
export EMAIL_NOTIFIY=xxxxxx

#### DB POSTGRES
export DB_NAME=gologs
export DB_HOST=localhost
export DB_USER=gologs
export DB_PASSWORD=1234
export DEBUG=true
export DEBUG_REQ=100000
./gologs
