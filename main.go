package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/karuppiah7890/publish-to-nats/pkg/config"
	"github.com/nats-io/nats.go"
)

// TODO: Write tests for all of this

var version string

func main() {
	log.Printf("version: %v", version)
	c, err := config.NewConfigFromEnvVars()
	if err != nil {
		log.Fatalf("error occurred while getting configuration from environment variables: %v", err)
	}

	// Connect to a server
	nc, err := nats.Connect(c.GetNatsServerUrl())
	if err != nil {
		log.Fatalf("error occurred while connecting to NATS server: %v", err)
	}

	messages, err := readJson(c.GetNatsMessagesJsonFilePath())
	if err != nil {
		log.Fatalf("error occurred while reading JSON to get messages to send to NATS server: %v", err)
	}

	for _, message := range messages {
		err := nc.Publish(message.Subject, []byte(message.Payload))
		if err != nil {
			log.Fatalf("error occurred while sending message to NATS server: %v", err)
		}
		fmt.Printf(".")
	}
}

type Message struct {
	Payload string `json:"payload"`
	Subject string `json:"subject"`
}

type Messages []Message

func readJson(filePath string) (Messages, error) {
	jsonFileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error occurred while reading JSON file at file path %s: %v", filePath, err)
	}

	var messages Messages

	err = json.Unmarshal(jsonFileContent, &messages)
	if err != nil {
		log.Fatalf("error occurred while parsing JSON file: %v", err)
	}

	return messages, nil
}
