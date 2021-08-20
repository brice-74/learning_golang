include .env
#export $(shell sed 's/=.*//' .env)

PROJECT_DIR="."
UPDOWN_FILE="create_x_table"
migrate-files:
	migrate create -seq -ext=.sql -dir=$(PROJECT_DIR)/migrations $(UPDOWN_FILE)
migrate-up:
	migrate -path=./migrations -database=$(GREENLIGHT_DB_DSN) up
migrate-down:
	migrate -path=./migrations -database=$(GREENLIGHT_DB_DSN) down
VERSION=1
migrate-version:
	migrate -path=./migrations -database=$(GREENLIGHT_DB_DSN) goto $(VERSION)
migrate-check-version:
	migrate -path=./migrations -database=$(GREENLIGHT_DB_DSN) version
migrate-force:
	migrate -path=./migrations -database$(GREENLIGHT_DB_DSN) force 1

start:
	go run ./cmd/api