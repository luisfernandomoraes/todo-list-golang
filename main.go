package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/database"
	envs "github.com/luisfernandomoraes/todo-list-golang/infraestructure/envs"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/logs"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/server"
)

func main() {
	envs.CheckEnvironmentVariables()
	logs.Info("Starting todo-list-golang service")
	database.StartDB()
	s := server.NewServer()
	s.Run()
}
