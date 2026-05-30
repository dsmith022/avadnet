package api

import (
	"avadnet/backend/internal/api/handlers"
	"avadnet/backend/internal/config"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API routes under /api/v1.
// Handlers may be unimplemented stubs returning 501; the register endpoint is implemented.
func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api/v1")

	// Auth
	auth := api.Group("/auth")
	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/login", handlers.LoginUser)

	// Missions
	missions := api.Group("/missions")
	missions.POST("", handlers.CreateMissionOrganization)
	missions.GET("", handlers.ListMissionOrganizations)
	missions.GET(":id", handlers.GetMissionOrganization)
	missions.PUT(":id", handlers.UpdateMissionOrganization)
	missions.DELETE(":id", handlers.DeleteMissionOrganization)
	missions.POST(":id/members", handlers.AddMissionMember)
	missions.GET(":id/members", handlers.ListMissionMembers)
	missions.DELETE(":id/members/:user_id", handlers.RemoveMissionMember)
	missions.POST(":id/ventures", handlers.AddMissionVenture)
	missions.GET(":id/ventures", handlers.ListMissionVentures)
	missions.DELETE(":id/ventures/:venture_id", handlers.RemoveMissionVenture)

	// Churches
	churches := api.Group("/churches")
	churches.POST("", handlers.CreateChurch)
	churches.GET("", handlers.ListChurches)
	churches.GET(":id", handlers.GetChurch)
	churches.PUT(":id", handlers.UpdateChurch)
	churches.DELETE(":id", handlers.DeleteChurch)
	churches.POST(":id/members", handlers.AddChurchMember)
	churches.GET(":id/members", handlers.ListChurchMembers)
	churches.PUT(":id/members/:user_id", handlers.UpdateChurchMemberRole)
	churches.DELETE(":id/members/:user_id", handlers.RemoveChurchMember)
}
