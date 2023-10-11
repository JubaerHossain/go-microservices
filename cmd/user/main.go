package main

import (
	"github.com/JubaerHossain/go-service/docs"
	"github.com/JubaerHossain/go-service/internal/user/config"
	"github.com/JubaerHossain/go-service/internal/user/routes"
	"github.com/JubaerHossain/gomd/gomd"
)

func main() {

	gomd.New() // Initialize gomd

	config.Register() // Register config

	routes.Register() // Register routes

	// kafka.Initialize() // Initialize Kafka consumer

	gomd.NoSqlConnection() // Initialize MongoDB connection

	config.Boot() // Boot the application

	docs.DocRegister() // Register config

	gomd.Run()
}
