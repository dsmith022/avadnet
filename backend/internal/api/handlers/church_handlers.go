package handlers

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// Church Creation Request
type CreateChurchRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"required"`
}

// Church Update Request
type UpdateChurchRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"required"`
}

// Get Church Request
type GetChurchRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

// List Churches Request
type ListChurchesRequest struct {
	Limit  int `form:"limit,default=10" binding:"min=1,max=100"`
	Offset int `form:"offset,default=0" binding:"min=0"`
}

// Delete Church Request
type DeleteChurchRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

// Church Members Requests
type AddChurchMemberRequest struct {
	UserID   uuid.UUID `json:"user_id" binding:"required"`
	ChurchID uuid.UUID `json:"church_id" binding:"required"`
	Role     string    `json:"role" binding:"required,oneof=admin member pastor"`
}

type RemoveChurchMemberRequest struct {
	ChurchID uuid.UUID `uri:"church_id" binding:"required"`
	UserID   uuid.UUID `uri:"user_id" binding:"required"`
}

type ListChurchMembersRequest struct {
	ChurchID uuid.UUID `uri:"church_id" binding:"required"`
	Limit    int       `form:"limit,default=10" binding:"min=1,max=100"`
	Offset   int       `form:"offset,default=0" binding:"min=0"`
}

type UpdateChurchMemberRoleRequest struct {
	ChurchID uuid.UUID `uri:"church_id" binding:"required"`
	UserID   uuid.UUID `uri:"user_id" binding:"required"`
	Role     string    `json:"role" binding:"required,oneof=admin member pastor"`
}

// Church Response
type ChurchResponse struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

// Church Member Response
type ChurchMemberResponse struct {
	UserID    uuid.UUID        `json:"user_id"`
	ChurchID  uuid.UUID        `json:"church_id"`
	Role      string           `json:"role"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	CreatedAt pgtype.Timestamp `json:"created_at,omitempty"`
}

// List Churches Response
type ListChurchesResponse struct {
	Churches []ChurchResponse `json:"churches"`
	Total    int              `json:"total"`
	Limit    int              `json:"limit"`
	Offset   int              `json:"offset"`
}

// List Church Members Response
type ListChurchMembersResponse struct {
	Members []ChurchMemberResponse `json:"members"`
	Total   int                    `json:"total"`
	Limit   int                    `json:"limit"`
	Offset  int                    `json:"offset"`
}
