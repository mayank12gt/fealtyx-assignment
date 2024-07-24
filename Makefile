include .env

.Phony: api/run
api/run:
	go run ./cmd/api 

