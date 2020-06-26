package client_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/walmartdigital/gomock-tutorial-code/pkg/client"
	"github.com/walmartdigital/gomock-tutorial-code/pkg/mocks"
)

var ctrl *gomock.Controller

func TestAll(t *testing.T) {
	ctrl = gomock.NewController(t)
	defer ctrl.Finish()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Client tests")
}

var _ = Describe("Read message", func() {
	var (
		fakeHTTPRequest *mocks.MockHTTPRequest
	)

	BeforeEach(func() {
		fakeHTTPRequest = mocks.NewMockHTTPRequest(ctrl)
	})

	It("should read a message from the server", func() {
		msg := client.ReadMessage(fakeHTTPRequest, "dogs")
		Expect(msg).To(Equal("Hi there, I love dogs!"))
	})
})
