package routes

import (
	c "github.com/JubaerHossain/go-service/internal/tenant/controllers"
	"github.com/JubaerHossain/gomd/config"
	"github.com/JubaerHossain/gomd/gomd"
	. "github.com/JubaerHossain/gomd/gomd"
	"github.com/gin-gonic/gin"
)

func Register() {
	Router.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"app": config.Config.GetString("App.Name"),
		}
		Res.Code(200).
			Message("success").
			Data(data).Json(c)
	})

	v1 := gomd.Router.Group("api/v1")
	v1.GET("tenants", c.TenantIndex())
	v1.POST("tenants", c.TenantCreate())
	v1.GET("tenants/:tenantId", c.TenantShow())
	v1.PUT("tenants/:tenantId", c.TenantUpdate())
	v1.DELETE("tenants/:tenantId", c.TenantDelete())
}
