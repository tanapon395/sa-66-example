package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/sa-66-example/entity"
)

// GET /genders
func ListGenders(c *gin.Context) {
	var genders []entity.Gender

	entity.DB().Find(&genders)
	c.JSON(http.StatusOK, genders)
}
