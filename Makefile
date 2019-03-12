# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on

### GMAIL SMTP
export GMAIL_PASSWORD=xxxxxx
export GMAIL_USER=xxxxxx
export EMAIL_NOTIFIY=xxxxxx

#### DB POSTGRES
export DB_NAME=xxxxxx
export DB_HOST=localhost
export DB_USER=xxxxxx
export DB_PASSWORD=xxxxxx

all:
	echo $$GOPATH
	##go get -d
	##go run *.go

test:
	@echo "\033[0;33m################################## go tests ##################################\033[0m"

push:
	@echo "--------------------------------------------------------------------------------------"
	@echo "\033[0;33m################################## build Docker gologs ##################################\033[0m"
	@docker build --no-cache -f Dockerfile -t your-regis/gologs .

	@echo "\033[0;33m################################## Login Docker ##################################\033[0m"
	@docker login
	@docker push your-regis/gologs:latest
	@echo "\033[0;32mGenerated\033[0m \033[0;33m[ok]\033[0m \n"

build:
	@echo "\033[0;33m################################## build prod exec: gologs ##################################\033[0m"
	@GOOS=linux go build -o gologs -ldflags="-s -w" && ./gologs
	@echo "\033[0;32mGenerated\033[0m \033[0;33m[ok]\033[0m \n"

brute:
	@echo "\033[0;33m################################## build prod exec: gologs ##################################\033[0m"
	@GOOS=linux go build -o gologs -ldflags="-s -w"
	@upx --brute gologs
	./gologs
	@echo "\033[0;32mGenerated\033[0m \033[0;33m[ok]\033[0m \n"
