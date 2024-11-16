include .env
export $(shell sed 's/=.*//' .env)

CURRENT_DIR=$(shell pwd)
PDB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

mig-up:
	migrate -path db/migrations -database '${PDB_URL}' -verbose up

mig-down:
	migrate -path db/migrations -database '${PDB_URL}' -verbose down

mig-force:
	migrate -path db/migrations -database '${PDB_URL}' -verbose force 1

create_mig:
	@echo "Enter file name: "; \
	read filename; \
	migrate create -ext sql -dir db/migrations -seq $$filename

swag:
	~/go/bin/swag init -g ./api/router.go -o api/docs
	
run:
	go run cmd/main.go

tidy:
	go mod tidy

test:
	go test ./storage/postgres

swag-change: 
	go get -u github.com/swaggo/swag
