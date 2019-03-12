#!/bin/bash
GOOS=linux go build -o gologs -ldflags="-s -w"

### GMAIL SMTP
export GMAIL_PASSWORD=xxxxxx
export GMAIL_USER=xxxxxx
export EMAIL_NOTIFIY=xxxxxx

#### DB POSTGRES
export DB_NAME=xxxxxx
export DB_HOST=localhost
export DB_USER=xxxxxx
export DB_PASSWORD=xxxxxx
./gologs
