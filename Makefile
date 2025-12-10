.PHONY: help test bench fuzz vet lint run clean

help: ## Mostrar esta ayuda
	@echo "Comandos disponibles:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

test: ## Ejecutar tests
	go test -v ./...

bench: ## Ejecutar benchmarks
	go test -bench=. -benchmem ./...

fuzz: ## Ejecutar fuzzing
	go test -fuzz=. ./...

vet: ## Ejecutar go vet
	go vet ./...

lint: ## Ejecutar golangci-lint (si está instalado)
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint no está instalado. Instala con: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

run-fundamentals: ## Ejecutar ejemplos de fundamentos
	@echo "Ejecutando ejemplos de fundamentos..."
	@cd fundamentals/types_structs && go run main.go
	@cd fundamentals/interfaces && go run main.go
	@cd fundamentals/methods && go run main.go
	@cd fundamentals/collections && go run main.go
	@cd fundamentals/errors && go run main.go
	@cd fundamentals/packages && go run main.go

run-concurrency: ## Ejecutar ejemplos de concurrencia
	@echo "Ejecutando ejemplos de concurrencia..."
	@cd concurrency/goroutines && go run main.go
	@cd concurrency/channels && go run main.go
	@cd concurrency/context && go run main.go
	@cd concurrency/sync && go run main.go
	@cd concurrency/worker_pool && go run main.go
	@cd concurrency/pipeline && go run main.go

clean: ## Limpiar archivos generados
	find . -name "*.test" -delete
	find . -name "*.out" -delete
	go clean -cache

