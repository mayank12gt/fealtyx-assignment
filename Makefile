include .env

.Phony: api/run
api/run:
	go run ./cmd/api -Port ${PORT} -Ollama_URL ${OLLAMA_BASE_URL} 

