package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Placeholder handlers for church endpoints. These return 501 Not Implemented
// until business logic and persistence are implemented.

func CreateChurch(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "create church not implemented"})
}

func ListChurches(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "list churches not implemented"})
}

func GetChurch(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "get church not implemented"})
}

func UpdateChurch(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "update church not implemented"})
}

func DeleteChurch(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "delete church not implemented"})
}

func AddChurchMember(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "add church member not implemented"})
}

func ListChurchMembers(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "list church members not implemented"})
}

func UpdateChurchMemberRole(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "update church member role not implemented"})
}

func RemoveChurchMember(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "remove church member not implemented"})
}
