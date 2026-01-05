package entity

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestAllfiled(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	registration := Registration{
		Name:      "Jessada",
		Email:     "guy@gmail.com",
		StudentID: "B6512866",
		Age:       19,
	}

	ok, err := registration.Validate()

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
}

func TestName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	registration := Registration{
		Name:      "",
		Email:     "guy@gmail.com",
		StudentID: "B6512866",
		Age:       19,
	}

	ok, err := registration.Validate()

	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err.Error()).To(gomega.Equal("Name Is Required"))
}

func TestEmail(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	registration := Registration{
		Name:      "Jessada",
		Email:     "dsad",
		StudentID: "B6512866",
		Age:       19,
	}

	ok ,err := registration.Validate()

	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err.Error()).To(gomega.Equal("Required Type Email"))
}

func TestAge(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	registration := Registration{
		Name:      "Jessada",
		Email:     "guy@gmail.com",
		StudentID: "B6512866",
		Age:       1,
	}

	ok, err := registration.Validate()

	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err.Error()).To(gomega.Equal("Age must be 18-60"))
}