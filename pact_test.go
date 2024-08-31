package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	pact := &dsl.Pact{
		Consumer:                 "GoConsumerContractTesting",
		Provider:                 "GoProviderContractTesting",
		DisableToolValidityCheck: true,
		LogDir:                   "logs",
		PactDir:                  "pacts",
	}

	// Start the API in the background
	go startAPI()

	// Verify the Provider
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8080",
		PactURLs:        []string{"./pacts/chattemplateconsumer-chattemplateprovider.json"}, // Path to the local Pact file
	})

	assert.NoError(t, err)
}

func startAPI() {
	http.HandleFunc("/actions/bcd/chat-template", func(w http.ResponseWriter, r *http.Request) {
		resp, err := Handler()
		if err != nil {
			http.Error(w, err.Error(), resp.StatusCode)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Body))
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}
