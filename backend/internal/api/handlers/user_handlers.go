package handlers

import (
	"avadnet/backend/internal/api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RegisterUser handles user registration.
// It validates the incoming request and returns a created UserResponse.
func RegisterUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// NOTE: persistence is not implemented yet. We return a minimal created response.
	id := uuid.New()
	now := time.Now()

	resp := models.UserResponse{
		ID:        id,
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	c.JSON(http.StatusCreated, resp)
}

// LoginUser is a placeholder login handler returning 501 Not Implemented.
func LoginUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "login not implemented"})
}
