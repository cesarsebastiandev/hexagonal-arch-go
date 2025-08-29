package user

import (
	"net/http"

	"hexagonal-architecture/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateUser(c *gin.Context) {

	//Translate the request
	//Validations
	//Consume the service
	//Translate the response
	var userCreateParams domain.User

	if err := c.BindJSON(&userCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//=====================================================
	// insertedId, err := h.UserService.Create(userCreateParams)
	_, err := h.UserService.Create(userCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hooops!"})

	}
	//=====================================================

	// c.JSON(200, gin.H{"user_id": insertedId})
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
