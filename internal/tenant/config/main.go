package config

import (
	"github.com/JubaerHossain/gomd/config"
	tenant "github.com/JubaerHossain/go-service/internal/tenant/models"
)

func Register() {
	config.Config.AddConfig("App", new(AppConfig))
	config.Config.AddConfig("NoSql", new(MongoConfig))
	config.Config.Load()
}

func Boot() {
	tenant.TenantSetup()
}
