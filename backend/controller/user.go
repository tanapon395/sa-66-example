package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/tanapon395/sa-66-example/entity"
	"golang.org/x/crypto/bcrypt"
)

// POST /users
func CreateUser(c *gin.Context) {
	var user entity.User

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
	var gender entity.Gender
	entity.DB().First(&gender, user.GenderID)
	if gender.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "gender not found"})
		return
	}

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้าง User
	u := entity.User{
		StudentID: user.StudentID,
		FirstName: user.FirstName, // ตั้งค่าฟิลด์ FirstName
		LastName:  user.LastName,  // ตั้งค่าฟิลด์ LastName
		Email:     user.Email,     // ตั้งค่าฟิลด์ Email
		Phone:     user.Phone,     // ตั้งค่าฟิลด์ Phone
		Profile:   user.Profile,   // ตั้งค่าฟิลด์ Profile
		LinkIn:    user.LinkIn,
		GenderID:  user.GenderID,
		Gender:    gender, // โยงความสัมพันธ์กับ Entity Gender
	}

	// บันทึก
	if err := entity.DB().Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Created success", "data": u})
}

// GET /user/:id
func GetUser(c *gin.Context) {
	var user entity.User
	id := c.Param("id")
	entity.DB().Preload("Gender").First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GET /users
func ListUsers(c *gin.Context) {
	var users []entity.User
	entity.DB().Preload("Gender").Find(&users)
	c.JSON(http.StatusOK, users)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user entity.User
	entity.DB().First(&user, id)
	if user.ID != 0 {
		entity.DB().Delete(&user)
		c.JSON(http.StatusOK, gin.H{"message": "Deleted success"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
	}

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
