# Guía de Contribución

## Estructura del Proyecto

Este proyecto está organizado por temas. Cada directorio contiene ejemplos independientes que puedes ejecutar.

## Cómo Ejecutar Ejemplos

### Fundamentos
```bash
cd fundamentals/types_structs
go run main.go
```

### Concurrencia
```bash
cd concurrency/goroutines
go run main.go
```

### Testing
```bash
cd testing/unit
go test -v
```

## Agregar Nuevos Ejemplos

1. Crea un nuevo directorio bajo el tema correspondiente
2. Incluye un `main.go` con ejemplos ejecutables
3. Agrega comentarios detallados explicando los conceptos
4. Si es relevante, agrega tests

## Convenciones

- Usa nombres descriptivos
- Incluye comentarios explicativos
- Compara con Java cuando sea relevante (para desarrolladores que vienen de Java)
- Mantén los ejemplos simples pero completos

