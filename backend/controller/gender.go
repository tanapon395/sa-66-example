package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/sa-66-example/entity"
)

// GET /genders
func ListGenders(c *gin.Context) {
	var genders []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM genders").Scan(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": genders})
}
