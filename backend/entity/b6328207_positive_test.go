package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameCantBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	c := Customer{
		Name:       "", //ผิด
		Email:      "ddd@gmail.com",
		CustomerID: "L0394857",
	}

	ok, err := govalidator.ValidateStruct(c)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name cant Blank"))

}

func TestCustomerID(t *testing.T) {
	g := NewGomegaWithT(t)
	c := Customer{
		Name:       "fgdg",
		Email:      "ddd@gmail.com",
		CustomerID: "A0394857",
	}

	ok, err := govalidator.ValidateStruct(c)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Please Enter L or M or H and Follow by number 7 letter"))

}

func TestCurrect(t *testing.T) {
	g := NewGomegaWithT(t)
	c := Customer{
		Name:       "fdgsdsd",
		Email:      "ddd@gmail.com",
		CustomerID: "L0394857",
	}

	ok, err := govalidator.ValidateStruct(c)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Current"))

}
