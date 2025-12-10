package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/handler"
	"github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/infrastructure"
	"github.com/josediaz/go-mastery-lab/architecture/clean_arch_api/internal/usecase"
)

// ============================================================================
// MAIN - Punto de entrada de la aplicación
// ============================================================================
// Aquí se ensamblan todas las capas usando Dependency Injection
// ============================================================================

func main() {
	// 1. Crear repositorio (infrastructure)
	userRepo := infrastructure.NewMemoryUserRepository()

	// 2. Crear casos de uso (usecase)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// 3. Crear handlers (handler)
	userHandler := handler.NewUserHandler(userUsecase)

	// 4. Configurar router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 5. Definir rutas
	r.Post("/users", userHandler.CreateUser)
	r.Get("/users/{id}", userHandler.GetUser)

	// 6. Iniciar servidor
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

