package controllers

import (
	"net/http"
	"time"

	"github.com/Eduardo-Brehm/mobile-api-go/internal/db"
	"github.com/Eduardo-Brehm/mobile-api-go/internal/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest is the request body for the register endpoint
type RegisterRequest struct {
	Tipo  string `json:"tipo"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

// LoginRequest is the request body for the login endpoint
type LoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type AuthController struct {
	userRepo *db.UserRepository
}

func NewAuthController(userRepo *db.UserRepository) *AuthController {
	return &AuthController{userRepo: userRepo}
}

func (ac *AuthController) Register(c echo.Context) error {

	//read the request body
	var req RegisterRequest

	//validate the request body
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "dados inválidos"})
	}

	//crypt the password
	hashedPassword, err := ac.userRepo.HashPassword(req.Senha)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "erro ao processar senha"})
	}

	//call ac.userRepo.CreateUser to create the user in the database
	createdUser, err := ac.userRepo.CreateUser(&db.User{
		Nome:         &req.Nome,
		Email:        req.Email,
		SenhaHash:    hashedPassword,
		CriadoEm:     time.Now(),
		AtualizadoEm: time.Now(),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "erro ao criar usuário"})
	}

	//return json with the created user
	return c.JSON(http.StatusCreated, createdUser)
}

func (ac *AuthController) Login(c echo.Context) error {
	//receive email and password from the request body
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// call ac.userRepo.GetUserByEmail to get the user from the database
	user, err := ac.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "email ou senha incorretos"})
	}

	// check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.SenhaHash), []byte(req.Senha))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "email ou senha incorretos"})
	}

	//if the password is correct, generate de JWT token and return it in the response
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "erro ao gerar token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
