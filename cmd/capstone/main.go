package main

import (
	"log"
	"os"

	"github.com/lucdoe/capstone/internal"
	"github.com/lucdoe/capstone/internal/app"
	"github.com/lucdoe/capstone/internal/app/databases"
	"gopkg.in/yaml.v3"
)

func LoadConfig(f string) (*internal.Config, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var config internal.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	log.Printf("Loaded config: %v", config)
	return &config, nil
}

func main() {
	databases.InitializeRedis()

	config, err := LoadConfig("gateway_config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	APIGatewayAPP, err := app.APIGatewayAPP(config)
	if err != nil {
		log.Fatal(err)
	}

	APIGatewayAPP.Router.Run(":8080")
}
