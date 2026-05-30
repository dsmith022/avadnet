package handlers

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// User Creation Request
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// User Login Request
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// User Update Request
type UpdateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
}

// Get User Request
type GetUserRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

// List Users Request
type ListUsersRequest struct {
	Limit  int `form:"limit,default=10" binding:"min=1,max=100"`
	Offset int `form:"offset,default=0" binding:"min=0"`
}

// Delete User Request
type DeleteUserRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

// User Response
type UserResponse struct {
	ID        uuid.UUID        `json:"id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

// User Login Response (with auth token)
type UserLoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token,omitempty"`
}

// List Users Response
type ListUsersResponse struct {
	Users  []UserResponse `json:"users"`
	Total  int            `json:"total"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
}
