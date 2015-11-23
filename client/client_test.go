package client_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"code.google.com/p/go-uuid/uuid"

	"github.com/tscolari/tshirts/challenge"
	"github.com/tscolari/tshirts/client"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var server *httptest.Server
	var serverHandler http.HandlerFunc

	var tClient *client.Client
	var authToken string
	var baseUrl string

	var request *http.Request

	BeforeEach(func() {
		request = nil
		authToken = uuid.New()
	})

	AfterEach(func() {
		server.Close()
	})

	JustBeforeEach(func() {
		server = httptest.NewServer(http.HandlerFunc(serverHandler))

		baseUrl = server.URL
		tClient = client.New(baseUrl, authToken)
	})

	Describe("FetchInks", func() {
		var inks challenge.Inks

		BeforeEach(func() {
			serverHandler = func(w http.ResponseWriter, r *http.Request) {
				request = r
				fmt.Fprintf(w, `{
					"inks": [
					{
						"id": "VN1348",
						"color": "#17B0D8",
						"cost": 12.39
					}
					]}`)
			}
		})

		JustBeforeEach(func() {
			var err error
			inks, err = tClient.FetchInks()
			Expect(err).ToNot(HaveOccurred())
		})

		It("sends an GET request", func() {
			Expect(request.Method).To(Equal("GET"))
		})

		It("sets the correct auth-token header", func() {
			Expect(request.Header.Get("Auth-Token")).To(Equal(authToken))
		})

		It("requests the correct path", func() {
			Expect(request.URL.Path).To(Equal("/inks"))
		})

		It("parses the response correctly", func() {
			Expect(inks[0].Color).To(Equal("#17B0D8"))
			Expect(inks[0].ID).To(Equal("VN1348"))
			Expect(inks[0].Cost).To(Equal(12.39))
		})
	})

	Describe("FetchQuestion", func() {
		var scenario challenge.Scenario

		BeforeEach(func() {
			serverHandler = func(w http.ResponseWriter, r *http.Request) {
				request = r
				fmt.Fprintf(w, `{
					"scenario_id": "9dc14adc-1d84-49af-a404-54ef47b0e3d4",
					"questions": [
					{
						"layers": [
						{
							"color": "#C68069",
							"volume": 6.450718063498126
						}
						]
					}
					]
				}`)
			}
		})

		JustBeforeEach(func() {
			var err error
			scenario, err = tClient.FetchQuestion()
			Expect(err).ToNot(HaveOccurred())
		})

		It("sends an GET request", func() {
			Expect(request.Method).To(Equal("GET"))
		})

		It("sets the correct auth-token header", func() {
			Expect(request.Header.Get("Auth-Token")).To(Equal(authToken))
		})

		It("requests the correct path", func() {
			Expect(request.URL.Path).To(Equal("/question/practice"))
		})

		It("parses the response correctly", func() {
			Expect(scenario.ID).To(Equal("9dc14adc-1d84-49af-a404-54ef47b0e3d4"))
			Expect(len(scenario.Questions)).To(Equal(1))
			Expect(len(scenario.Questions[0].Layers)).To(Equal(1))
			Expect(scenario.Questions[0].Layers[0].Color).To(Equal("#C68069"))
			Expect(scenario.Questions[0].Layers[0].Volume).To(Equal(6.450718063498126))
		})
	})

	Describe("PostAnswer", func() {
		BeforeEach(func() {
			serverHandler = func(w http.ResponseWriter, r *http.Request) {
				request = r
			}
		})

		JustBeforeEach(func() {
			var err error
			_, err = tClient.PostAnswer(challenge.Solution{})
			Expect(err).ToNot(HaveOccurred())
		})

		It("sends an POST request", func() {
			Expect(request.Method).To(Equal("POST"))
		})

		It("sets the correct auth-token header", func() {
			Expect(request.Header.Get("Auth-Token")).To(Equal(authToken))
		})

		It("requests the correct path", func() {
			Expect(request.URL.Path).To(Equal("/answer/practice"))
		})
	})
})
