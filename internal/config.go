package internal

import (
	"encoding/json"
	"log"
	"os"

	"kvm-manager/types"
)

var Config *types.Config

func InitProvisioning() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config.json"
	}
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	cfg := &types.Config{}
	if err := json.Unmarshal(f, cfg); err != nil {
		log.Fatalf("Invalid config JSON: %v", err)
	}
	Config = cfg
	log.Println("Configuration loaded.")
}
