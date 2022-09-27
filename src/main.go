package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/database"
	envs "github.com/luisfernandomoraes/todo-list-golang/infraestructure/envs"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/logger"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/server"
)

func main() {
	envs.CheckEnvironmentVariables()
	logger.GetLogger().Info().Str("use-case", "infra").Msg("Starting todo-list-golang service")
	database.StartDB()
	s := server.NewServer()
	s.Run()
}
