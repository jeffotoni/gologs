#!/bin/bash

echo "--------------------------------------------------------------------------------------"
echo "\033[0;33m################################## build docker DigitalOcean ##################################\033[0m"

docker stop gologs
docker rm gologs
sleep 1
docker run -d -p 22335:22335 -p 22334:22334 -e DB_NAME=gologs -e DB_HOST=localhost -e DB_USER=gologs -e DB_PASSWORD=1234 -e DB_PORT=5432 -e DB_SSL=disable -e DB_SORCE=postgres -e GMAIL_PASSWORD=love2020graf -e GMAIL_USER="" -e EMAIL_NOTIFIY="" --rm --name gologs jeffotoni/gologs

