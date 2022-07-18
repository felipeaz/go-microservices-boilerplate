package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"microservices-boilerplate/internal/serviceB/service"
	assertion "microservices-boilerplate/test/assertion/serviceB"
	mock "microservices-boilerplate/test/mocks/pkg"
)

func TestServiceB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service B Suits")
}

var _ = Describe("Service B", func() {
	logMock := new(mock.Logger)
	s := service.New(logMock)

	Context("Testing CRUD Operations", func() {

		Context("Getting All items", func() {
			When("Request succeeds", func() {
				expectedItems := assertion.ItemArray
				It("Should return all items from DB", func() {
					resp, err := s.GetAll()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItems))
				})
			})
		})

		Context("Getting one item by ID", func() {
			When("Request succeeds", func() {
				expectedItem := assertion.NewItemWithID(assertion.SampleID.String())
				It("Should return an item with given ID", func() {
					resp, err := s.GetOneByID(expectedItem.ID)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItem))
				})
			})
		})

		Context("Creating an item", func() {
			When("Request succeeds", func() {
				itemInput := assertion.NewItemWithoutID()
				It("Should return the created object", func() {
					resp, err := s.Create(itemInput)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp.ID).NotTo(BeNil())
				})
			})
		})

		Context("Updating an item", func() {
			When("Request succeeds", func() {
				inputItem := assertion.NewItemWithID(assertion.SampleID.String())
				It("Should return nothing", func() {
					err := s.Update(assertion.SampleID, inputItem)
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Context("Deleting an item", func() {
			When("Request succeeds", func() {
				It("Should return nothing", func() {
					err := s.Delete(assertion.SampleID)
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})
})
