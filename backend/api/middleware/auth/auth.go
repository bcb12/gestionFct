package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bcb12/gestionFct/internal/helpers"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"golang.org/x/crypto/bcrypt"
)

// User is the structure which holds one user from the database.
type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name,omitempty"`
	FirstSurname string    `json:"first_surname,omitempty"`
	Password     string    `json:"-"`
	Active       int       `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

const dbTimeout = time.Second * 3

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization header is missing")
			//http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			helpers.ErrorJSON(w, errors.New("authorization header is missing"), http.StatusUnauthorized)
			return
		}

		splits := strings.Split(authHeader, " ")
		if len(splits) != 2 || strings.ToLower(splits[0]) != "bearer" {
			log.Println("Invalid Authorization header format")
			//http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			helpers.ErrorJSON(w, errors.New("invalid Authorization header format"), http.StatusUnauthorized)
			return
		}

		// Verificar el token (este paso puede variar dependiendo de tu implementación de autenticación)
		token := splits[1]
		if !isValidToken(token) { // La función isValidToken se tiene que definir según tu lógica de validación de tokens
			log.Println("Invalid token")
			//http.Error(w, "Invalid token", http.StatusUnauthorized)
			helpers.ErrorJSON(w, errors.New("invalid token"), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// GetByEmail returns one user by email
func GetByEmail(db *sql.DB, email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, first_surname, password, user_active, created_at, updated_at from users where email = $1`

	var user User
	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.FirstSurname,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func isValidToken(token string) bool {
	isValid := false

	return isValid
}

// Create user

// Check permissions
