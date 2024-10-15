package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Struct Definitions

type ProducerResponseWrapper struct {
	Data ProducerResponseData `json:"data"`
}

type ProducerResponseData struct {
	Type       string                     `json:"type"`
	Attributes ProducerResponseAttributes `json:"attributes"`
}

type ProducerResponseAttributes struct {
	ProducerChatTemplates []ProducerChatTemplate `json:"templates"`
}

type ProducerChatTemplate struct {
	Id             string   `json:"id"`
	Consumer       string   `json:"consumer"`
	Title          string   `json:"title"`
	Category       string   `json:"category"`
	Message        string   `json:"message"`
	Queues         []string `json:"queues" dynamodbav:"-"`
	QueueList      []Queue  `json:"-" dynamodbav:"queues"`    // Ignored in JSON
	RecordDisabled bool     `json:"recordDisabled,omitempty"` // Omitted if false
}

type Queue struct {
	QueueName string `json:"queueName"`
	QueueARN  string `json:"queueArn"`
}

// ValidConsumers defines the set of allowed consumer values
var ValidConsumers = map[string][]ProducerChatTemplate{
	"fullserve": {
		{Id: "1", Consumer: "fullserve", Title: "title1", Category: "category1", Message: "message1", Queues: []string{"queue1", "queue2"}},
		{Id: "2", Consumer: "fullserve", Title: "title2", Category: "category2", Message: "message2", Queues: []string{"queue3", "queue4"}},
		{Id: "3", Consumer: "fullserve", Title: "title3", Category: "category3", Message: "message2", Queues: []string{"queue5", "queue6", "queue7"}},
	},
	"veripark": {
		{Id: "4", Consumer: "veripark", Title: "title4", Category: "category4", Message: "message4", Queues: []string{"queue8"}},
		{Id: "5", Consumer: "veripark", Title: "title5", Category: "category5", Message: "message5", Queues: []string{"queue9", "queue10"}},
	},
	"salesforce": {
		{Id: "6", Consumer: "salesforce", Title: "title6", Category: "category6", Message: "message6", Queues: []string{"queue6", "queue7", "queue8"}},
		{Id: "7", Consumer: "salesforce", Title: "title7", Category: "category7", Message: "message7", Queues: []string{"queue9", "queue10", "queue11"}},
		{Id: "8", Consumer: "salesforce", Title: "title8", Category: "category8", Message: "message8", Queues: []string{"queue12", "queue1", "queue3"}},
		{Id: "9", Consumer: "salesforce", Title: "title9", Category: "category9", Message: "message9", Queues: []string{"queue4", "queue6", "queue9"}},
	},
}

// filterByConsumer returns the chat templates for a valid consumer
func filterByConsumer(consumer string) ([]ProducerChatTemplate, bool) {
	templates, exists := ValidConsumers[consumer]
	return templates, exists
}

// Handler processes the HTTP request and returns the appropriate response
func Handler(consumer string) (events.APIGatewayProxyResponse, error) {
	// Validate the consumer
	templates, valid := filterByConsumer(consumer)
	if !valid {
		// Return 404 if consumer is invalid
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "Consumer not found"}`,
		}, nil
	}

	// Create the response
	response := ProducerResponseWrapper{
		Data: ProducerResponseData{
			Type: "ChatTemplate",
			Attributes: ProducerResponseAttributes{
				ProducerChatTemplates: templates,
			},
		},
	}

	// Marshal the response to JSON
	responseBody, err := json.Marshal(response)
	if err != nil {
		// Return 500 if JSON marshaling fails
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "Internal Server Error"}`,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}

// main sets up the HTTP server and routes
func main() {
	http.HandleFunc("/actions/bcd/chat-template", func(w http.ResponseWriter, r *http.Request) {
		// Extract the 'consumer' query parameter
		consumer := r.URL.Query().Get("consumer")

		// Check if 'consumer' is provided
		if consumer == "" {
			http.Error(w, `{"error": "Consumer parameter is required"}`, http.StatusNotFound)
			return
		}

		// Call the handler with the provided consumer
		resp, err := Handler(consumer)
		if err != nil {
			http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
			return
		}

		// Set headers and write the response
		for key, value := range resp.Headers {
			w.Header().Set(key, value)
		}
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Body))
	})

	// Start the HTTP server on port 8080
	log.Println("Producer API server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
