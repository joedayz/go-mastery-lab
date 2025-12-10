# Cross-Compiling en Go

Go soporta cross-compiling nativamente. No necesitas un entorno de compilación específico para cada plataforma.

## Compilar para diferentes plataformas

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o app-linux main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o app-windows.exe main.go

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o app-macos-intel main.go

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o app-macos-arm main.go
```

## Variables de entorno

- `GOOS`: Sistema operativo (linux, windows, darwin, etc.)
- `GOARCH`: Arquitectura (amd64, arm64, 386, etc.)

## Ejemplo

```bash
# Compilar para todas las plataformas comunes
./build-all.sh
```

