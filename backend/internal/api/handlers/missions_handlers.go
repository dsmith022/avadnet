package handlers

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// Mission Organization Requests
type CreateMissionOrganizationRequest struct {
	Name        string                     `json:"name" binding:"required"`
	Description string                     `json:"description" binding:"required"`
	Members     []AddMissionMemberRequest  `json:"members" binding:"required,dive"`
	Ventures    []AddMissionVentureRequest `json:"ventures" binding:"dive"`
}

type UpdateMissionOrganizationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type GetMissionOrganizationRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

type ListMissionOrganizationsRequest struct {
	Limit  int `form:"limit,default=10" binding:"min=1,max=100"`
	Offset int `form:"offset,default=0" binding:"min=0"`
}

type DeleteMissionOrganizationRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

// Mission Members Requests
type AddMissionMemberRequest struct {
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	MissionID uuid.UUID `json:"mission_id" binding:"required"`
	Role      string    `json:"role" binding:"required,oneof=admin member"`
}

type RemoveMissionMemberRequest struct {
	MissionID uuid.UUID `json:"mission_id" binding:"required"`
	UserID    uuid.UUID `json:"user_id" binding:"required"`
}

type ListMissionMembersRequest struct {
	MissionID uuid.UUID `uri:"mission_id" binding:"required"`
	Limit     int       `form:"limit,default=10" binding:"min=1,max=100"`
	Offset    int       `form:"offset,default=0" binding:"min=0"`
}

// Mission Ventures Requests
type AddMissionVentureRequest struct {
	VentureID   uuid.UUID                 `json:"venture_id" binding:"required"`
	MissionID   uuid.UUID                 `json:"mission_id" binding:"required"`
	Description string                    `json:"description" binding:"required"`
	Members     []AddMissionMemberRequest `json:"members" binding:"dive"`
}

type RemoveMissionVentureRequest struct {
	MissionID uuid.UUID `json:"mission_id" binding:"required"`
	VentureID uuid.UUID `json:"venture_id" binding:"required"`
}

type ListMissionVenturesRequest struct {
	MissionID uuid.UUID `uri:"mission_id" binding:"required"`
	Limit     int       `form:"limit,default=10" binding:"min=1,max=100"`
	Offset    int       `form:"offset,default=0" binding:"min=0"`
}

// Mission Organization Response
type MissionOrganizationResponse struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

// Mission Member Response
type MissionMemberResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	MissionID uuid.UUID `json:"mission_id"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}

// Mission Venture Response
type MissionVentureResponse struct {
	ID          uuid.UUID `json:"id"`
	MissionID   uuid.UUID `json:"mission_id"`
	VentureID   uuid.UUID `json:"venture_id"`
	VentureName string    `json:"venture_name"`
	Description string    `json:"description"`
}
