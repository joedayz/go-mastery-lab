# Go Mastery Lab 🚀

Laboratorio completo de dominio avanzado de Go para desarrolladores que buscan alcanzar nivel Senior.

## 📋 Tabla de Contenidos

### 🎯 Guías de Inicio
- **[LEARNING_PATH.md](LEARNING_PATH.md)** ⭐ **EMPIEZA AQUÍ** - Orden recomendado paso a paso
- **[INSTALLATION.md](INSTALLATION.md)** - Cómo instalar Go
- **[GETTING_STARTED.md](GETTING_STARTED.md)** - Conceptos clave para desarrolladores Java

### 📚 Contenido del Curso
1. [Fundamentos Avanzados](#1-fundamentos-avanzados-del-lenguaje)
2. [Concurrencia](#2-concurrencia)
3. [Estructura de Proyectos y Clean Architecture](#3-estructura-de-proyectos-y-clean-architecture)
4. [Networking y APIs](#4-networking-y-apis)
5. [Persistencia](#5-persistencia)
6. [Go Modules y Builds](#6-go-modules-builds-y-optimización)
7. [Testing Avanzado](#7-testing-avanzado)
8. [Herramientas del Ecosistema](#8-herramientas-del-ecosistema)
9. [Docker, CI/CD y Despliegue](#9-docker-cicd-y-despliegue)
10. [Patrones y Prácticas](#10-patrones-y-prácticas)

## 🎯 Objetivo

Este repositorio está diseñado para desarrolladores que vienen de otros lenguajes (especialmente Java) y necesitan dominar Go a nivel Senior. Cada sección incluye:

- ✅ Ejemplos prácticos y ejecutables
- ✅ Comentarios detallados explicando conceptos
- ✅ Comparaciones con otros lenguajes cuando es relevante
- ✅ Mejores prácticas y patrones comunes
- ✅ Ejercicios y demos sugeridas

## 🏗️ Estructura del Proyecto

```
go-mastery-lab/
├── fundamentals/          # Fundamentos avanzados
│   ├── types_structs/    # Tipos y structs
│   ├── interfaces/       # Interfaces implícitas
│   ├── methods/          # Métodos y receptores
│   ├── collections/      # Slices, maps, arrays
│   └── errors/           # Manejo de errores
├── concurrency/          # Concurrencia
│   ├── goroutines/       # Goroutines básicas
│   ├── channels/         # Channels y select
│   ├── context/          # Context con cancelación
│   ├── sync/             # Mutex, WaitGroup, etc.
│   ├── worker_pool/      # Worker pools
│   └── pipeline/         # Pipelines y Fan-In/Out
├── architecture/         # Arquitectura limpia
│   └── clean_arch_api/   # API REST con Clean Architecture
├── http/                 # Networking
│   ├── rest_api/         # API REST
│   ├── grpc_service/     # Servicio gRPC
│   └── websockets/       # WebSockets
├── persistence/          # Persistencia
│   └── sqlc_demo/        # CRUD con sqlc
├── testing/              # Testing avanzado
│   ├── unit/             # Unit tests
│   ├── fuzz/             # Fuzzing
│   └── benchmarks/       # Benchmarks
├── cli/                  # CLI y builds
│   ├── build_flags/      # Build tags
│   └── cross_compile/    # Cross-compiling
├── profiling/            # Profiling
│   └── pprof_demo/       # pprof examples
├── patterns/             # Patrones
│   ├── functional_options/ # Opciones funcionales
│   └── retry_backoff/    # Retry y circuit breaker
└── docker/               # Docker y CI/CD
    └── ci_cd/            # GitHub Actions
```

## 🚀 Cómo Usar Este Repositorio

### ⭐ **EMPIEZA AQUÍ**: Lee primero `LEARNING_PATH.md` para el orden recomendado

1. **Sigue la ruta de aprendizaje**: `LEARNING_PATH.md` tiene el orden paso a paso
2. **Lee los comentarios**: Cada archivo tiene explicaciones detalladas
3. **Ejecuta los ejemplos**: `go run` en cada directorio
4. **Experimenta**: Modifica el código y observa los resultados

### Ejecutar Ejemplos

```bash
# Fundamentos
cd fundamentals/types_structs && go run main.go

# Concurrencia
cd concurrency/goroutines && go run main.go

# Testing
cd testing/unit && go test -v

# Usar Makefile
make run-fundamentals
make run-concurrency
make test
make bench
```

### Requisitos

- Go 1.21 o superior
- Para algunos ejemplos necesitarás instalar dependencias:
  ```bash
  go mod download
  ```

## 📚 Requisitos Previos

- Go 1.21 o superior ([Guía de Instalación](INSTALLATION.md))
- Conocimiento básico de programación
- Familiaridad con conceptos de programación orientada a objetos (viniendo de Java)

## 🎓 Notas para Desarrolladores Java

Este repositorio incluye notas especiales para desarrolladores que vienen de Java:

- **Interfaces**: En Go son implícitas (duck typing), no necesitas `implements`
- **Herencia**: No existe, usa composición
- **Generics**: Disponibles desde Go 1.18 (similar a Java generics)
- **Concurrencia**: Muy diferente a Java threads, más ligera y eficiente
- **Gestión de memoria**: Automática como Java, pero sin JVM

## 📝 Licencia

Apache License 2.0

## 🤝 Contribuciones

Este es un proyecto de aprendizaje. Siéntete libre de:
- Agregar más ejemplos
- Mejorar documentación
- Corregir errores
- Sugerir mejoras

---

**¡Bienvenido al mundo de Go! 🐹**

