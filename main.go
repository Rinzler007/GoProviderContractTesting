package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

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
	QueueList      []Queue  `json:"-" dynamodbav:"queues"`
	RecordDisabled bool     `json:"recordDisabled"`
}

type Queue struct {
	QueueName string `json:"queueName"`
	QueueARN  string `json:"queueArn"`
}

func Handler() (events.APIGatewayProxyResponse, error) {
	response := ProducerResponseWrapper{
		Data: ProducerResponseData{
			Type: "ChatTemplate",
			Attributes: ProducerResponseAttributes{
				ProducerChatTemplates: []ProducerChatTemplate{
					{
						Id:       "1",
						Consumer: "fullserve",
						Title:    "title1",
						Category: "category1",
						Message:  "message1",
						Queues:   []string{"queue1", "queue2"},
					},
					{
						Id:       "2",
						Consumer: "fullserve",
						Title:    "title2",
						Category: "category2",
						Message:  "message2",
						Queues:   []string{"queue3", "queue4"},
					},
					{
						Id:       "3",
						Consumer: "fullserve",
						Title:    "title3",
						Category: "category3",
						Message:  "message3",
						Queues:   []string{"queue5"},
					},
					{
						Id:       "4",
						Consumer: "fullserve",
						Title:    "title4",
						Category: "category4",
						Message:  "message4",
						Queues:   []string{"queue6", "queue7", "queue8"},
					},
				},
			},
		},
	}
	responseBody, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}

func main() {
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
