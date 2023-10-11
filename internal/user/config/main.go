package config

import (
	"github.com/JubaerHossain/gomd/config"
	user "github.com/JubaerHossain/go-service/internal/user/models"
)

func Register() {
	config.Config.AddConfig("App", new(AppConfig))
	config.Config.AddConfig("NoSql", new(MongoConfig))
	config.Config.Load()
}

func Boot() {
	user.UserSetup()
}
