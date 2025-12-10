package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3" // Driver SQLite (necesitas instalarlo)
)

// ============================================================================
// PERSISTENCIA CON DATABASE/SQL
// ============================================================================
// database/sql es el paquete estándar para trabajar con bases de datos
// Similar a JDBC en Java, pero más simple
// ============================================================================

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func main() {
	// Abrir conexión (pool de conexiones automático)
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Configurar pool de conexiones
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Crear tabla
	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	// Insertar usuario
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, err := insertUser(ctx, db, "John Doe", "john@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted user with ID: %d\n", userID)

	// Obtener usuario
	user, err := getUser(ctx, db, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved user: %+v\n", user)

	// Obtener múltiples usuarios
	users, err := getUsers(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total users: %d\n", len(users))
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	return err
}

func insertUser(ctx context.Context, db *sql.DB, name, email string) (int64, error) {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	result, err := db.ExecContext(ctx, query, name, email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func getUser(ctx context.Context, db *sql.DB, id int64) (*User, error) {
	query := `SELECT id, name, email, created_at FROM users WHERE id = ?`
	row := db.QueryRowContext(ctx, query, id)

	var user User
	var createdAt string
	err := row.Scan(&user.ID, &user.Name, &user.Email, &createdAt)
	if err != nil {
		return nil, err
	}

	user.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	return &user, nil
}

func getUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	query := `SELECT id, name, email, created_at FROM users`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var createdAt string
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &createdAt); err != nil {
			return nil, err
		}
		user.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		users = append(users, user)
	}
	return users, rows.Err()
}

