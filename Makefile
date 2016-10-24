export VERSION:=$(shell go run power.go version)
include JWT

dev:
	docker-compose up linux expose

release:
	dobi -v release:push


watch:
	dobi -v watch


