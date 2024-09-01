package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	// Start the API in the background
	go startAPI()

	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/pacts", dir)
	var logDir = fmt.Sprintf("%s/logs", dir)
	pact := &dsl.Pact{
		Provider:                 "GoProviderContractTesting",
		DisableToolValidityCheck: true,
		LogDir:                   logDir,
		PactDir:                  pactDir,
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://localhost:8080",      //provider's URL
		BrokerURL:                  "https://harsh.pactflow.io/", //link to your remote Contract broker
		BrokerToken:                "2_KfMXbOMRXKAd30PwopTg",     //your PactFlow token
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	// Verify the Provider
	log.Println("[debug] start verification")
	_, err = pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8080",
		PactURLs:        []string{"pacts/goconsumercontracttesting-goprovidercontracttesting.json"}, // Path to the local Pact file
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
}

func startAPI() {
	main()
}
