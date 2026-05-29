package handlers

import "github.com/google/uuid"

type AddMissionMemberRequest struct {
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	MissionID uuid.UUID `json:"mission_id" binding:"required"`
	Role      string    `json:"role" binding:"required,oneof=admin member"`
}

type AddMissionVentureRequest struct {
	VentureID   uuid.UUID                 `json:"venture_id" binding:"required"`
	MissionID   uuid.UUID                 `json:"mission_id" binding:"required"`
	Description string                    `json:"description" binding:"required"`
	Members     []AddMissionMemberRequest `json:"members" binding:"dive"`
}

type CreateMissionOrganizationRequest struct {
	Name        string                     `json:"name" binding:"required"`
	Description string                     `json:"description" binding:"required"`
	Members     []AddMissionMemberRequest  `json:"members" binding:"required,dive"`
	Ventures    []AddMissionVentureRequest `json:"ventures" binding:"dive"`
}
