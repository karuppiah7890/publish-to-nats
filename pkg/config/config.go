package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/nats-io/nats.go"
)

// All configuration is through environment variables

const NATS_SERVER_URL_ENV_VAR = "NATS_SERVER_URL"
const DEFAULT_NATS_SERVER_URL = nats.DefaultURL
const NATS_MESSAGES_JSON_FILE_PATH_ENV_VAR = "NATS_MESSAGES_JSON_FILE_PATH"
const DFAULT_NATS_MESSAGES_JSON_FILE_PATH = "messages.json"

type Config struct {
	natsServerUrl            string
	natsMessagesJsonFilePath string
}

func NewConfigFromEnvVars() (*Config, error) {
	natsServerUrl, err := getNatsServerUrl()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting nats server url: %v", err)
	}

	natsMessagesJsonFilePath, err := getNatsMessagesJsonFilePath()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting nats messages json file path: %v", err)
	}

	return &Config{
		natsServerUrl:            natsServerUrl,
		natsMessagesJsonFilePath: natsMessagesJsonFilePath,
	}, nil
}

// Get nats server url
func getNatsServerUrl() (string, error) {
	natsServerUrl, ok := os.LookupEnv(NATS_SERVER_URL_ENV_VAR)
	if !ok {
		natsServerUrl = DEFAULT_NATS_SERVER_URL
	}

	return natsServerUrl, nil
}

// Get nats messages json file path
func getNatsMessagesJsonFilePath() (string, error) {
	natsMessagesJsonFilePath, ok := os.LookupEnv(NATS_MESSAGES_JSON_FILE_PATH_ENV_VAR)
	if !ok {
		natsMessagesJsonFilePath = DFAULT_NATS_MESSAGES_JSON_FILE_PATH
	}

	_, err := os.Stat(natsMessagesJsonFilePath)

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return "", fmt.Errorf("NATS messages JSON file does not exist at path %s", natsMessagesJsonFilePath)
		}

		return "", fmt.Errorf("could not find file info of the NATS messages JSON file at path %s: %v", natsMessagesJsonFilePath, err)
	}

	return natsMessagesJsonFilePath, nil
}

func (c *Config) GetNatsServerUrl() string {
	return c.natsServerUrl
}

func (c *Config) GetNatsMessagesJsonFilePath() string {
	return c.natsMessagesJsonFilePath
}
