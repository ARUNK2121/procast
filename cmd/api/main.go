package main

import (
	"log"

	"github.com/ARUNK2121/procast/cmd/api/docs"
	"github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/di"
	"github.com/ARUNK2121/procast/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	docs.SwaggerInfo.Title = "Procast"
	docs.SwaggerInfo.Description = utils.Description
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "procast.jerseyhubmadebyarunk.store"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
