# 🗺️ Ruta de Aprendizaje - Orden Recomendado

Esta guía te lleva paso a paso por todo el contenido del proyecto en el orden óptimo para aprender Go desde cero hasta nivel Senior.

## 📋 Fase 0: Preparación (15 minutos)

### 1. Instalar Go
- ✅ Lee: `INSTALLATION.md`
- ✅ Verifica: `go version`
- ✅ Prueba: Crea un programa "Hello World"

### 2. Entender la Estructura del Proyecto
- ✅ Lee: `README.md` (visión general)
- ✅ Lee: `GETTING_STARTED.md` (conceptos clave si vienes de Java)

---

## 📚 Fase 1: Fundamentos Avanzados (2-3 días)

**Objetivo**: Dominar los conceptos básicos pero fundamentales de Go.

### Día 1: Tipos y Estructuras

1. **Tipos y Structs**
   ```bash
   cd fundamentals/types_structs
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: structs, embedding, métodos básicos
   - **Ejercicio**: Crea tu propio struct con métodos

2. **Interfaces Implícitas**
   ```bash
   cd fundamentals/interfaces
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: duck typing, type assertions
   - **Ejercicio**: Agrega un nuevo proveedor de pagos

3. **Métodos y Receptores**
   ```bash
   cd fundamentals/methods
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: value vs pointer receivers
   - **Ejercicio**: Modifica métodos y observa diferencias

### Día 2: Colecciones y Errores

4. **Slices, Maps, Arrays**
   ```bash
   cd fundamentals/collections
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: len, cap, make, append
   - **Ejercicio**: Crea tu propio sistema de inventario

5. **Manejo de Errores**
   ```bash
   cd fundamentals/errors
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: errors.Is(), errors.As(), errors.Join(), wrapping
   - **Ejercicio**: Crea funciones con manejo de errores robusto

6. **Paquetes Importantes**
   ```bash
   cd fundamentals/packages
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: context, io, fmt, sync, reflect
   - **Ejercicio**: Experimenta con cada paquete

---

## ⚡ Fase 2: Concurrencia (3-4 días) ⭐ **MUY IMPORTANTE**

**Objetivo**: Dominar la concurrencia, el superpoder de Go.

### Día 3: Goroutines y Channels

7. **Goroutines**
   ```bash
   cd concurrency/goroutines
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: qué son goroutines, WaitGroup, closures
   - **Ejercicio**: Crea múltiples goroutines que procesen datos

8. **Channels**
   ```bash
   cd concurrency/channels
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: buffered vs unbuffered, select, range
   - **Ejercicio**: Implementa comunicación entre goroutines

### Día 4: Context y Sincronización

9. **Context**
   ```bash
   cd concurrency/context
   go run main.go
   ```
   - Lee los comentarios en `main.go`
   - Entiende: cancelación, timeouts, valores en context
   - **Ejercicio**: Crea operaciones con timeouts

10. **Sync Primitives**
    ```bash
    cd concurrency/sync
    go run main.go
    ```
    - Lee los comentarios en `main.go`
    - Entiende: Mutex, RWMutex, WaitGroup, Once, Cond, Pool
    - **Ejercicio**: Crea estructuras thread-safe

### Día 5: Patrones Avanzados

11. **Worker Pools**
    ```bash
    cd concurrency/worker_pool
    go run main.go
    ```
    - Lee los comentarios en `main.go`
    - Entiende: procesar miles de tareas eficientemente
    - **Ejercicio**: Crea tu propio worker pool

12. **Pipelines**
    ```bash
    cd concurrency/pipeline
    go run main.go
    ```
    - Lee los comentarios en `main.go`
    - Entiende: pipelines, fan-out, fan-in
    - **Ejercicio**: Crea un pipeline de transformación de datos

---

## 🏗️ Fase 3: Arquitectura y Estructura (1-2 días)

**Objetivo**: Aprender a estructurar proyectos profesionales.

13. **Clean Architecture**
    ```bash
    cd architecture/clean_arch_api
    # Lee primero: README.md
    # Luego explora los archivos en orden:
    # 1. internal/domain/user.go
    # 2. internal/repository/user_repository.go
    # 3. internal/usecase/user_usecase.go
    # 4. internal/handler/user_handler.go
    # 5. internal/infrastructure/memory_user_repository.go
    # 6. cmd/api/main.go
    
    # Ejecutar (si tienes las dependencias):
    go run cmd/api/main.go
    ```
    - Entiende: separación de capas, inversión de dependencias
    - **Ejercicio**: Agrega un nuevo endpoint siguiendo la arquitectura

---

## 🌐 Fase 4: Networking (1 día)

14. **REST API**
    ```bash
    cd http/rest_api
    go run main.go
    # En otra terminal:
    curl http://localhost:8080/users
    ```
    - Entiende: net/http, routers, middlewares
    - **Ejercicio**: Agrega más endpoints y middlewares

---

## 💾 Fase 5: Persistencia (1 día)

15. **Database/SQL**
    ```bash
    cd persistence/sql_demo
    # Nota: Necesitas instalar driver SQLite primero:
    # go get github.com/mattn/go-sqlite3
    go run main.go
    ```
    - Entiende: database/sql, contextos, pooling
    - **Ejercicio**: Crea operaciones CRUD completas

---

## 🧪 Fase 6: Testing (1 día)

16. **Unit Tests**
    ```bash
    cd testing/unit
    go test -v
    ```
    - Lee los comentarios en `main_test.go`
    - Entiende: table-driven tests, mocks
    - **Ejercicio**: Escribe tests para tus propios ejemplos

17. **Benchmarks**
    ```bash
    cd testing/benchmarks
    go test -bench=.
    ```
    - Entiende: cómo medir rendimiento
    - **Ejercicio**: Compara diferentes implementaciones

18. **Fuzzing**
    ```bash
    cd testing/fuzz
    go test -fuzz=FuzzReverse
    ```
    - Entiende: encontrar bugs con inputs aleatorios
    - **Ejercicio**: Crea tests de fuzzing para funciones propias

---

## 🎨 Fase 7: Patrones (1 día)

19. **Functional Options**
    ```bash
    cd patterns/functional_options
    go run main.go
    ```
    - Entiende: patrón común en Go para configuración
    - **Ejercicio**: Crea tu propia función con opciones

20. **Retry y Backoff**
    ```bash
    cd patterns/retry_backoff
    go run main.go
    ```
    - Entiende: reintentar operaciones fallidas
    - **Ejercicio**: Implementa retry para llamadas HTTP

21. **Circuit Breaker**
    ```bash
    cd patterns/circuit_breaker
    go run main.go
    ```
    - Entiende: prevenir llamadas a servicios caídos
    - **Ejercicio**: Integra circuit breaker en una API

---

## 🛠️ Fase 8: Herramientas y Builds (1 día)

22. **Build Tags y Cross-Compiling**
    ```bash
    cd cli/cross_compile
    # Lee: README.md
    ./build-all.sh
    ```
    - Entiende: compilar para múltiples plataformas
    - **Ejercicio**: Compila tu proyecto para Linux/Windows/macOS

23. **Profiling con pprof**
    ```bash
    cd profiling/pprof_demo
    go run main.go
    # En otra terminal:
    go tool pprof http://localhost:6060/debug/pprof/profile
    ```
    - Lee: `README.md` en el directorio
    - Entiende: cómo encontrar cuellos de botella
    - **Ejercicio**: Profiliza tu propio código

24. **Docker y CI/CD**
    ```bash
    cd docker
    # Lee los Dockerfiles
    docker build -f Dockerfile -t go-app .
    ```
    - Entiende: multi-stage builds, imágenes mínimas
    - **Ejercicio**: Crea Dockerfile para tu proyecto

---

## 📊 Resumen de Tiempo Estimado

| Fase | Tiempo | Prioridad |
|------|--------|-----------|
| Fase 0: Preparación | 15 min | ⭐⭐⭐ |
| Fase 1: Fundamentos | 2-3 días | ⭐⭐⭐ |
| Fase 2: Concurrencia | 3-4 días | ⭐⭐⭐⭐⭐ |
| Fase 3: Arquitectura | 1-2 días | ⭐⭐⭐⭐ |
| Fase 4: Networking | 1 día | ⭐⭐⭐ |
| Fase 5: Persistencia | 1 día | ⭐⭐⭐ |
| Fase 6: Testing | 1 día | ⭐⭐⭐⭐ |
| Fase 7: Patrones | 1 día | ⭐⭐⭐ |
| Fase 8: Herramientas | 1 día | ⭐⭐⭐ |

**Total estimado**: 10-14 días de estudio intensivo

---

## 🎯 Ruta Rápida (Si tienes prisa)

Si necesitas aprender rápido, sigue este orden mínimo:

1. ✅ `INSTALLATION.md` - Instalar Go
2. ✅ `fundamentals/types_structs` - Structs básicos
3. ✅ `fundamentals/interfaces` - Interfaces
4. ✅ `fundamentals/collections` - Slices y maps
5. ✅ `fundamentals/errors` - Manejo de errores
6. ✅ `concurrency/goroutines` - Goroutines
7. ✅ `concurrency/channels` - Channels
8. ✅ `concurrency/context` - Context
9. ✅ `concurrency/sync` - Sync primitives
10. ✅ `architecture/clean_arch_api` - Arquitectura
11. ✅ `http/rest_api` - REST API
12. ✅ `testing/unit` - Testing básico

**Tiempo mínimo**: 5-7 días

---

## 💡 Consejos de Estudio

1. **No te saltes pasos**: Cada fase construye sobre la anterior
2. **Ejecuta el código**: No solo leas, ejecuta y modifica
3. **Experimenta**: Cambia valores, rompe cosas, aprende de errores
4. **Toma notas**: Escribe tus propias observaciones
5. **Practica**: Después de cada sección, crea algo propio
6. **Revisa**: Vuelve a leer código anterior cuando aprendas algo nuevo

---

## ✅ Checklist de Progreso

Marca cada fase cuando la completes:

- [ ] Fase 0: Preparación
- [ ] Fase 1: Fundamentos Avanzados
- [ ] Fase 2: Concurrencia
- [ ] Fase 3: Arquitectura
- [ ] Fase 4: Networking
- [ ] Fase 5: Persistencia
- [ ] Fase 6: Testing
- [ ] Fase 7: Patrones
- [ ] Fase 8: Herramientas

---

## 🚀 Siguiente Paso

**Empieza ahora mismo**:

```bash
# 1. Verifica que Go esté instalado
go version

# 2. Ve al primer ejemplo
cd fundamentals/types_structs

# 3. Lee y ejecuta
cat main.go  # Lee el código
go run main.go  # Ejecuta

# 4. Modifica y experimenta
# Abre main.go en tu editor y haz cambios
```

---

**¡Buena suerte en tu aprendizaje! 🎓**

