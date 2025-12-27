# 🚀 Guía de Inicio Rápido - Go Mastery Lab

**Bienvenido a Go Mastery Lab!** Esta guía te llevará desde cero hasta ejecutar tu primer código en menos de 10 minutos.

## ⚡ Instalación Rápida (5 minutos)

### macOS
```bash
# Opción 1: Homebrew (Recomendado)
brew install go

# Opción 2: Instalador oficial
# Descarga desde https://go.dev/dl/
```

### Linux
```bash
# Opción 1: Gestor de paquetes
sudo apt install golang-go  # Ubuntu/Debian
sudo dnf install golang      # Fedora/RHEL

# Opción 2: Instalador oficial
wget https://go.dev/dl/go1.21.x.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.x.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### Windows
1. Descarga el instalador `.msi` desde https://go.dev/dl/
2. Ejecuta el instalador (configura PATH automáticamente)

### ✅ Verificar Instalación
```bash
go version
# Debería mostrar: go version go1.21.x ...
```

## 🎯 Conceptos Clave (Si Vienes de Java)

### 1. No hay Clases → Usa Structs
```go
// Java: public class User { private String name; }
// Go:
type User struct {
    name string
}
```

### 2. Interfaces son Implícitas
```go
// Java: class Dog implements Animal
// Go: Si Dog tiene los métodos de Animal, automáticamente lo implementa
type Animal interface {
    Speak() string
}
type Dog struct{}
func (d Dog) Speak() string { return "Woof" }
// Dog automáticamente implementa Animal ✅
```

### 3. No hay Excepciones → Errores como Valores
```go
// Java: throw new Exception("error")
// Go:
result, err := doSomething()
if err != nil {
    return err
}
```

### 4. Concurrencia es Diferente
```go
// Java: new Thread(() -> {...}).start()
// Go: go func() { ... }()  ← Mucho más ligero
go processData()
```

## 🏃 Tu Primer Programa (2 minutos)

```bash
# 1. Crear directorio
mkdir hello-go
cd hello-go

# 2. Inicializar módulo
go mod init hello-go

# 3. Crear main.go
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
EOF

# 4. Ejecutar
go run main.go
# Salida: Hello, Go!
```

## 📚 Orden Recomendado de Estudio

### Fase 1: Fundamentos (2-3 días)
1. **Tipos y Structs** → `fundamentals/types_structs/`
2. **Interfaces** → `fundamentals/interfaces/`
3. **Métodos** → `fundamentals/methods/`
4. **Colecciones** → `fundamentals/collections/`
5. **Errores** → `fundamentals/errors/`

### Fase 2: Concurrencia ⭐ **MUY IMPORTANTE** (3-4 días)
6. **Goroutines** → `concurrency/goroutines/`
7. **Channels** → `concurrency/channels/`
8. **Select** → `concurrency/channels/` (incluido)
9. **Context** → `concurrency/context/`
10. **Sync** → `concurrency/sync/`
11. **Worker Pools** → `concurrency/worker_pool/`
12. **Pipelines** → `concurrency/pipeline/`

### Fase 3: Arquitectura y APIs (2-3 días)
13. **Clean Architecture** → `architecture/clean_arch_api/`
14. **REST API** → `http/rest_api/`

### Fase 4: Testing y Más (1-2 días)
15. **Unit Tests** → `testing/unit/`
16. **Patrones** → `patterns/`

## 🎓 Cómo Usar Este Repositorio

### Paso 1: Ejecutar Ejemplos
```bash
# Ir a un ejemplo
cd fundamentals/types_structs

# Leer el código
cat main.go

# Ejecutar
go run main.go
```

### Paso 2: Experimentar
- Modifica el código
- Cambia valores
- Agrega nuevas funciones
- Observa qué pasa

### Paso 3: Practicar
- Crea tu propio código basado en los ejemplos
- Resuelve problemas similares
- Lee los comentarios en cada archivo

## 📖 Guía Detallada

Para una ruta de aprendizaje completa paso a paso con ejercicios y tiempos estimados, lee:
- **[LEARNING_PATH.md](LEARNING_PATH.md)** - Ruta completa detallada

## 🆘 Solución de Problemas Comunes

### "go: command not found"
```bash
# Verificar instalación
which go  # macOS/Linux
where go  # Windows

# Si falta, agregar a PATH:
export PATH=$PATH:/usr/local/go/bin  # macOS/Linux
```

### Versión incorrecta
```bash
# Verificar versión
go version

# Limpiar cache si hay problemas
go clean -cache
```

## ✅ Checklist de Inicio

- [ ] Go instalado (`go version` funciona)
- [ ] Primer programa ejecutado exitosamente
- [ ] Entendido conceptos básicos (structs, interfaces, errores)
- [ ] Listo para empezar con fundamentos

## 🚀 Siguiente Paso

**Empieza ahora mismo:**

```bash
# 1. Ve al primer ejemplo
cd fundamentals/types_structs

# 2. Lee y ejecuta
go run main.go

# 3. Modifica y experimenta
# Abre main.go en tu editor
```

---

**¿Listo para dominar Go?** Sigue la ruta en `LEARNING_PATH.md` para una guía paso a paso completa.

**¡Buena suerte! 🎓**

