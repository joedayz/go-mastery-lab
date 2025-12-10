# Clean Architecture en Go

Este ejemplo muestra cómo estructurar un proyecto Go siguiendo Clean Architecture.

## Estructura del Proyecto

```
clean_arch_api/
├── cmd/
│   └── api/
│       └── main.go          # Punto de entrada
├── internal/
│   ├── domain/              # Entidades y reglas de negocio
│   ├── repository/          # Interfaces de repositorio
│   ├── usecase/             # Casos de uso
│   ├── handler/             # HTTP handlers
│   └── infrastructure/      # Implementaciones concretas
├── pkg/                     # Paquetes reutilizables
└── go.mod
```

## Capas

1. **Domain**: Entidades puras, sin dependencias externas
2. **Repository**: Interfaces para acceso a datos
3. **Usecase**: Lógica de negocio
4. **Handler**: HTTP handlers (presentación)
5. **Infrastructure**: Implementaciones concretas (DB, HTTP client, etc.)

## Principios

- **Dependency Inversion**: Las capas internas no dependen de las externas
- **Separation of Concerns**: Cada capa tiene una responsabilidad clara
- **Testability**: Fácil de testear con mocks

