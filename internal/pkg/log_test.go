package pkg_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"microservices-boilerplate/internal/pkg"
	assertion "microservices-boilerplate/internal/test/assertion/pkg"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suits")
}

var _ = Describe("Log", func() {
	Context("Generating application logs", func() {
		Context("When Debug is disabled", func() {
			logger := pkg.NewLogger(assertion.LogTime, false)
			_, fileErr := os.Open(assertion.LogFile)
			When("Logging Info type", func() {
				It("Should call info.Println", func() {
					logger.Info(assertion.InfoLogMessage)

					Expect(fileErr).ShouldNot(HaveOccurred())
					Expect(assertion.LogFile).To(BeAnExistingFile())
				})
			})
			When("Logging Warn type", func() {
				It("Should call warn.Println", func() {
					logger.Warn(assertion.WarnLogMessage)

					Expect(fileErr).ShouldNot(HaveOccurred())
					Expect(assertion.LogFile).To(BeAnExistingFile())
				})
			})
			When("Logging Error type", func() {
				It("Should call info.Println", func() {
					logger.Error(assertion.ErrLogMessage)

					Expect(fileErr).ShouldNot(HaveOccurred())
					Expect(assertion.LogFile).To(BeAnExistingFile())
				})
			})
			When("Logging Debug type", func() {
				It("Should not call debug.Println", func() {
					logger.Debug(assertion.DebugLogMessage)

					Expect(fileErr).ShouldNot(HaveOccurred())
					Expect(assertion.LogFile).To(BeAnExistingFile())
				})
			})
		})
		Context("When Debug is enabled", func() {
			logger := pkg.NewLogger(assertion.LogTime, true)
			_, fileErr := os.Open(assertion.LogFile)
			When("Logging Debug type", func() {
				It("Should call debug.Println", func() {
					logger.Debug(assertion.DebugLogMessage)

					Expect(fileErr).ShouldNot(HaveOccurred())
					Expect(assertion.LogFile).To(BeAnExistingFile())
				})
			})
		})
	})
})

var _ = AfterSuite(func() {
	err := os.Remove(assertion.LogFile)

	Expect(err).NotTo(HaveOccurred())
})
