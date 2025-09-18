version = latest

all:run

.PHONY: run
run:
	go run cmd/server/main.go

# 编译
# GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o task main.go
.PHONY: build
build:
	go build -o docero cmd/server/main.go

.PHONY: docker
docker:
	sh ./script/build.sh ${version}