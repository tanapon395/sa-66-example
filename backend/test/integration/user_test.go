package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tanapon395/sa-66-example/controller"
	"github.com/tanapon395/sa-66-example/entity"
)

type resp struct {
	Error string `json:"error"`
}

func TestCreateUser(t *testing.T) {

	t.Run(`created user success`, func(t *testing.T) {
		r := gin.Default()
		r.POST("/users", controller.CreateUser)
		user := entity.User{
			StudentID: "B5000000",
			FirstName: "unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}
		jsonValue, _ := json.Marshal(user)
		reqFound, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, reqFound)

		fmt.Println(w.Body)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run(`create user is error as field "student_id" is not match pattern`, func(t *testing.T) {
		r := gin.Default()
		r.POST("/users", controller.CreateUser)
		user := entity.User{
			StudentID: "X5003049", // ผิด
			FirstName: "unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}
		jsonValue, _ := json.Marshal(user)
		reqFound, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, reqFound)
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		var respJson resp
		json.Unmarshal(body, &respJson)

		// สื่งที่คาดหวังจากการทดสอบ
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", user.StudentID), respJson.Error)

	})

}
