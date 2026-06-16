# Mobile API (Go)

**This is a study project.** It was built to practice backend development with Go, REST APIs, JWT authentication, and MySQL.

A simple REST API for user authentication and profile management.

## Tech stack

- Go + [Echo](https://echo.labstack.com/)
- MySQL
- JWT authentication

## Getting started

1. Copy `.env.example` to `.env` and fill in your database and JWT settings.
2. Run the server:

```bash
go run cmd/main.go
```

The API runs on `http://localhost:3000`.

## Authentication

Protected routes require a JWT token in the header:

```
Authorization: Bearer <token>
```

Tokens expire after 24 hours.

## Endpoints

Base path: `/api/v1`

### Auth

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/autenticacao/cadastro` | No | Register a new user |
| POST | `/autenticacao/login` | No | Login and receive a JWT token |

**Register** — request body:

```json
{
  "tipo": "user",
  "nome": "John Doe",
  "email": "john@example.com",
  "senha": "password123"
}
```

**Login** — request body:

```json
{
  "email": "john@example.com",
  "senha": "password123"
}
```

Response:

```json
{
  "token": "<jwt>"
}
```

### Profile

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/perfil/eu` | Yes | Get the authenticated user's profile |
| PATCH | `/perfil/eu` | Yes | Update the authenticated user's profile |
| GET | `/perfil/:username` | No | Get a public profile by username |

**Update profile** — request body:

```json
{
  "nomeDeUsuario": "johndoe",
  "bio": "Hello!",
  "fotoUrl": "https://example.com/photo.jpg",
  "siteUrl": "https://example.com"
}
```

All fields are optional except what you want to change. `bio`, `fotoUrl`, and `siteUrl` can be `null`.
