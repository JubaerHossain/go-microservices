package main

import (
	"github.com/JubaerHossain/gomd/gomd"
	"github.com/JubaerHossain/go-service/internal/tenant/config"
	"github.com/JubaerHossain/go-service/docs"
	// "github.com/JubaerHossain/go-service/internal/user/pkg/kafka"
	"github.com/JubaerHossain/go-service/internal/tenant/routes"
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
