package main

import (
	"flag"
	"log"

	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/handler"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/repository"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/server"
	"github.com/Yerlan-Tleubekov/real-time-forum/backend/internal/app/service"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.json", "path to config file")
}

func main() {
	flag.Parse()
	// sessionTokens := make(map[int]string)

	if err := server.ReadConfig(configPath); err != nil {
		log.Fatal(err)
	}

	config := server.GetConfig()

	db, err := repository.OpenDB(config.Database, config.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}

	sessionTokens := repository.NewSessionTokens()
	chatHubs := repository.NewChatHubs()

	repos := repository.NewRepository(db, sessionTokens, chatHubs)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	s := server.NewServer(handler.InitHandler())
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
