# Guía de Inicio Rápido

## 📦 Instalación de Go

**¿Primera vez instalando Go?** Lee la [Guía de Instalación Completa](INSTALLATION.md) que incluye:
- Instalación para macOS, Linux y Windows
- Configuración post-instalación
- Solución de problemas comunes
- Verificación de instalación

**Instalación rápida**:
- **macOS**: `brew install go` o descarga desde https://go.dev/dl/
- **Linux**: `sudo apt install golang-go` o descarga oficial
- **Windows**: Descarga el instalador MSI desde https://go.dev/dl/

## Para Desarrolladores que Vienen de Java

Si vienes de Java, estos son los conceptos clave que debes entender:

### 1. No hay Clases, hay Structs

```go
// Java
public class User {
    private String name;
    public User(String name) { this.name = name; }
}

// Go
type User struct {
    name string
}
func NewUser(name string) *User {
    return &User{name: name}
}
```

### 2. Interfaces son Implícitas

```go
// En Java necesitas: class Dog implements Animal
// En Go, si Dog tiene los métodos de Animal, automáticamente implementa Animal
type Animal interface {
    Speak() string
}
type Dog struct{}
func (d Dog) Speak() string { return "Woof" }
// Dog automáticamente implementa Animal
```

### 3. No hay Excepciones, hay Errores

```go
// Java: throw new Exception("error")
// Go: return nil, errors.New("error")

result, err := doSomething()
if err != nil {
    return err
}
```

### 4. Concurrencia es Diferente

```go
// Java: Thread thread = new Thread(() -> {...});
// Go: go func() { ... }()

go processData() // Muy ligero, no como threads
```

### 5. Gestión de Memoria

- Go tiene garbage collector como Java
- Pero no hay JVM
- Los binarios son estáticos y autocontenidos

## Orden Recomendado de Estudio

1. **Fundamentos** (1-2 días)
   - Tipos y structs
   - Interfaces
   - Métodos
   - Slices y maps
   - Errores

2. **Concurrencia** (2-3 días) ⭐ **MUY IMPORTANTE**
   - Goroutines
   - Channels
   - Select
   - Context
   - Sync primitives

3. **Arquitectura** (1-2 días)
   - Clean Architecture
   - Estructura de proyectos

4. **Networking** (1-2 días)
   - REST APIs
   - Middlewares

5. **Testing** (1 día)
   - Unit tests
   - Benchmarks
   - Fuzzing

6. **Patrones** (1 día)
   - Functional options
   - Retry/Backoff
   - Circuit breaker

## Recursos Adicionales

- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Blog](https://go.dev/blog/)

## Próximos Pasos

1. Ejecuta los ejemplos en orden
2. Modifica el código y experimenta
3. Intenta resolver problemas similares
4. Lee el código fuente de proyectos Go populares

