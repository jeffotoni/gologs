#!/bin/bash
GOOS=linux go build -ldflags="-s -w" -o worker-client-pg worker-client-pg.go

### GMAIL SMTP
export GMAIL_PASSWORD=xxxxxx
export GMAIL_USER=xxxxxx
export EMAIL_NOTIFIY=xxxxxx

#### DB POSTGRES
export DB_NAME=gologs
export DB_HOST=localhost
export DB_USER=gologs
export DB_PASSWORD=1234

export DEBUG=false
export DEBUG_REQ=500000

export MEMORY=200000
./worker-client-pg