package backend

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type BorrowBook struct {
	gorm.Model
	Email       string
	return_book string
}

func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	borrowbook := BorrowBook{
		Email:       "thanphirom@gmail.com",
		return_book: "วันที่ 9",
	}
	ok, err := govalidator.ValidateStruct(borrowbook)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	g.Expect(err.Error()).To(Equal(""))

}
