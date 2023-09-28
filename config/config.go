package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type WatcherConfig struct {
	Port                   string
	AlgorandProviderUrl    string
	UpdateTimeoutInSeconds int
}

const DEFAULT_PORT = "8080"
const DEFAULT_PROVIDER_URL = "https://testnet-api.algonode.cloud"
const DEFAULT_UPDATE_TIMEOUT = 10

func NewWatcherConfig() *WatcherConfig {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	port := os.Getenv("PORT")
	algorandProviderUrl := os.Getenv("ALGORAND_PROVIDER_URL")
	updateTimeoutInSecondsStr := os.Getenv("UPDATE_TIMEOUT_IN_SECONDS")

	if port == "" {
		port = DEFAULT_PORT
	}

	if algorandProviderUrl == "" {
		algorandProviderUrl = DEFAULT_PROVIDER_URL
	}

	updateTimeoutInSeconds, err := strconv.Atoi(updateTimeoutInSecondsStr)
	if err != nil {
		updateTimeoutInSeconds = DEFAULT_UPDATE_TIMEOUT
	}

	return &WatcherConfig{
		Port:                   port,
		AlgorandProviderUrl:    algorandProviderUrl,
		UpdateTimeoutInSeconds: updateTimeoutInSeconds,
	}
}
