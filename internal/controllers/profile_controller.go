package controllers

import (
	"net/http"

	"github.com/Eduardo-Brehm/mobile-api-go/internal/db"
	"github.com/labstack/echo/v4"
)

// UpdateProfileRequest is the payload for updating a profile
type UpdateProfileRequest struct {
	NomeDeUsuario string  `json:"nomeDeUsuario"`
	Bio           *string `json:"bio"`
	FotoUrl       *string `json:"fotoUrl"`
	SiteUrl       *string `json:"siteUrl"`
}

type ProfileController struct {
	profileRepo *db.ProfileRepository
	userRepo    *db.UserRepository
}

func NewProfileController(profileRepo *db.ProfileRepository, userRepo *db.UserRepository) *ProfileController {
	return &ProfileController{profileRepo: profileRepo, userRepo: userRepo}
}

// GetProfile retrieves the authenticated user's profile
func (pc *ProfileController) GetProfile(c echo.Context) error {
	// Get userID from middleware context
	userID := c.Get("userId").(string)

	// Get the profile from the database
	profile, err := pc.profileRepo.GetProfileByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "profile not found"})
	}

	return c.JSON(http.StatusOK, profile)
}

// UpdateProfile updates the authenticated user's profile
func (pc *ProfileController) UpdateProfile(c echo.Context) error {
	// Parse request body
	var req UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	// Get userID from middleware context
	userID := c.Get("userId").(string)

	// Get the current profile
	profile, err := pc.profileRepo.GetProfileByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "profile not found"})
	}

	// Update the profile fields
	profile.NomeDeUsuario = req.NomeDeUsuario
	profile.Bio = req.Bio
	profile.FotoUrl = req.FotoUrl
	profile.SiteUrl = req.SiteUrl

	// Save the updated profile
	updatedProfile, err := pc.profileRepo.UpdateProfile(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error updating profile"})
	}

	return c.JSON(http.StatusOK, updatedProfile)
}

// GetPublicProfile retrieves a public profile by username (no authentication required)
func (pc *ProfileController) GetPublicProfile(c echo.Context) error {
	// Get username from URL parameter
	username := c.Param("username")

	// Get the profile from the database
	profile, err := pc.profileRepo.GetProfileByUsername(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "profile not found"})
	}

	return c.JSON(http.StatusOK, profile)
}
