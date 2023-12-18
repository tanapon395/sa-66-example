package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/sa-66-example/entity"
)

// GET /genders
func ListGenders(c *gin.Context) {
	var genders []entity.Gender

	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.Find(&genders)
	c.JSON(http.StatusOK, genders)
}
