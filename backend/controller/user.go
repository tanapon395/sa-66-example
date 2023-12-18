package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/sa-66-example/entity"
	"golang.org/x/crypto/bcrypt"
)

// POST /users
func CreateUser(c *gin.Context) {
	var user entity.User
	var gender entity.Gender

	// bind เข้าตัวแปร user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	// ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", user.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// สร้าง User
	u := entity.User{
		Gender:    gender,               // โยงความสัมพันธ์กับ Entity Gender
		FirstName: user.FirstName,       // ตั้งค่าฟิลด์ FirstName
		LastName:  user.LastName,        // ตั้งค่าฟิลด์ LastName
		Email:     user.Email,           // ตั้งค่าฟิลด์ Email
		Password:  string(hashPassword), // เข้ารหัสผ่าน
		Phone:     user.Phone,           // ตั้งค่าฟิลด์ Phone
		Profile:   user.Profile,         // ตั้งค่าฟิลด์ Profile
	}

	// บันทึก
	if err := entity.DB().Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})
}

// GET /user/:id
func GetUser(c *gin.Context) {
	var user entity.User
	id := c.Param("id")
	if err := entity.DB().Preload("Gender").Raw("SELECT * FROM users WHERE id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users
func ListUsers(c *gin.Context) {
	var users []entity.User
	if err := entity.DB().Preload("Gender").Raw("SELECT * FROM users").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /users
func UpdateUser(c *gin.Context) {
	var user entity.User
	var result entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", user.ID).First(&result); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
