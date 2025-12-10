# Herramientas del Ecosistema Go

## go vet

Analizador estático que encuentra errores comunes:

```bash
go vet ./...
```

## golangci-lint

Linter completo que combina múltiples herramientas:

```bash
# Instalar
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Ejecutar
golangci-lint run

# Con configuración
golangci-lint run --config .golangci.yml
```

## gofmt

Formatea código automáticamente:

```bash
gofmt -w .  # Escribir cambios
gofmt -d .  # Mostrar diferencias
```

## goimports

Formatea y organiza imports:

```bash
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

## pprof

Profiling de CPU, memoria, goroutines:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

## race detector

Detecta data races:

```bash
go test -race ./...
go run -race main.go
```

## go mod

Gestión de dependencias:

```bash
go mod init          # Inicializar módulo
go mod tidy          # Limpiar dependencias
go mod download      # Descargar dependencias
go mod vendor        # Crear vendor directory
go mod verify        # Verificar checksums
```

## go work

Workspaces para múltiples módulos:

```bash
go work init
go work use ./module1 ./module2
```

## Delve (debugger)

Debugger para Go:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug main.go
```

