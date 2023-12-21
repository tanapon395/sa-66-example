package unit

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/tanapon395/sa-66-example/entity"
)

func TestStudentID(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`student_id is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: "", // ผิดตรงนี้
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})

	t.Run(`student_id pattern is not true`, func(t *testing.T) {
		user := entity.User{
			StudentID: "K5000000", // ผิดตรงนี้
			FirstName: "unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", user.StudentID)))
	})

	t.Run(`student_id is valid`, func(t *testing.T) {
		user := entity.User{
			StudentID: "B5000000",
			FirstName: "unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "", // ผิดตรงนี้
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestPhoneNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`phone_number is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: "B5000000",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "", // ผิดตรงนี้
			Profile:   "",
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone is required"))

	})

	t.Run(`phone_number check 10 digit`, func(t *testing.T) {
		user := entity.User{
			StudentID: "B5000000",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "080800000000", // ผิดตรงนี้ มี 11 ตัว
			Profile:   "",
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as stringlength(10|10)", user.Phone)))

	})
}
