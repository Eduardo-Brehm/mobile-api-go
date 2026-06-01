package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ProfileRepository struct {
	db *sql.DB // this is the database connection
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db} // this is the constructor
}

func (pr *ProfileRepository) CreateProfile(profile *Profile) (*Profile, error) {
	id := uuid.New().String()
	now := time.Now()
	profile.ID = id
	profile.CriadoEm = now
	profile.AtualizadoEm = now

	//Prepare the query
	stmt, err := pr.db.Prepare(
		"INSERT INTO perfis (id, usuarioId, tipo, nomeDeUsuario, bio, fotoUrl, siteUrl, criadoEm, atualizadoEm) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)

	// check for errors
	if err != nil {
		return nil, err
	}

	//Execute the query
	_, err = stmt.Exec(
		id,
		profile.UsuarioID,
		profile.Tipo,
		profile.NomeDeUsuario,
		profile.Bio,
		profile.FotoUrl,
		profile.SiteUrl,
		now,
		now,
	)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

// GetProfileByID searches for a profile by ID and returns the profile if found, otherwise returns an error
func (pr *ProfileRepository) GetProfileByUserID(userID string) (*Profile, error) {

	//Prepare the query
	var profile Profile

	//Execute the query
	err := pr.db.QueryRow(
		"SELECT id, usuarioId, tipo, nomeDeUsuario, bio, fotoUrl, siteUrl, criadoEm, atualizadoEm FROM perfis WHERE usuarioId = ?",
		userID,
	).Scan(&profile.ID, &profile.UsuarioID, &profile.Tipo, &profile.NomeDeUsuario, &profile.Bio, &profile.FotoUrl, &profile.SiteUrl, &profile.CriadoEm, &profile.AtualizadoEm)

	//Check for errors
	if err != nil {
		return nil, err
	}

	//Return the profile
	return &profile, nil
}

// GetProfileByID searches for a profile by ID and returns the profile if found, otherwise returns an error
func (pr *ProfileRepository) GetProfileByUsername(username string) (*Profile, error) {

	//Prepare the query
	var profile Profile

	//Execute the query
	err := pr.db.QueryRow(
		"SELECT id, usuarioId, tipo, nomeDeUsuario, bio, fotoUrl, siteUrl, criadoEm, atualizadoEm FROM perfis WHERE nomeDeUsuario = ?",
		username,
	).Scan(&profile.ID, &profile.UsuarioID, &profile.Tipo, &profile.NomeDeUsuario, &profile.Bio, &profile.FotoUrl, &profile.SiteUrl, &profile.CriadoEm, &profile.AtualizadoEm)

	//Check for errors
	if err != nil {
		return nil, err
	}

	//Return the profile
	return &profile, nil
}

func (pr *ProfileRepository) UpdateProfile(profile *Profile) (*Profile, error) {
	now := time.Now()
	profile.AtualizadoEm = now

	//Prepare the query
	stmt, err := pr.db.Prepare(
		"UPDATE perfis SET usuarioId = ?, tipo = ?, nomeDeUsuario = ?, bio = ?, fotoUrl = ?, siteUrl = ?, atualizadoEm = ? WHERE id = ?",
	)

	// check for errors
	if err != nil {
		return nil, err
	}

	//Execute the query
	_, err = stmt.Exec(
		profile.UsuarioID,
		profile.Tipo,
		profile.NomeDeUsuario,
		profile.Bio,
		profile.FotoUrl,
		profile.SiteUrl,
		now,
		profile.ID,
	)

	if err != nil {
		return nil, err
	}

	return profile, nil
}
