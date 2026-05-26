package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository is a repository for the User model
type UserRepository struct {
	db *sql.DB // this is the database connection
}

// NewUserRepository returns a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db} // this is the constructor
}

// GetUserByEmail searches for a user by email and returns the user if found, otherwise returns an error
func (ur *UserRepository) GetUserByEmail(email string) (*User, error) {

	//Prepare the query
	var user User

	//Execute the query
	err := ur.db.QueryRow(
		"SELECT id, nome, email, senhaHash, criadoEm, atualizadoEm FROM usuarios WHERE email = ?",
		email,
	).Scan(&user.ID, &user.Nome, &user.Email, &user.SenhaHash, &user.CriadoEm, &user.AtualizadoEm)

	//Check for errors
	if err != nil {
		return nil, err
	}

	//Return the user
	return &user, nil
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(user *User) (*User, error) {

	id := uuid.New().String()
	now := time.Now()
	user.ID = id
	user.CriadoEm = now
	user.AtualizadoEm = now

	//Prepare the query
	stmt, err := ur.db.Prepare(
		"INSERT INTO usuarios (id, nome, email, senhaHash, criadoEm, atualizadoEm) VALUES (?, ?, ?, ?, ?, ?)",
	)

	// check for errors
	if err != nil {
		return nil, err
	}

	//Execute the query
	_, err = stmt.Exec(
		id,
		user.Nome,
		user.Email,
		user.SenhaHash,
		now,
		now,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// HashPassword crypts the password using bcrypt and returns the hashed password
func (ur *UserRepository) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
