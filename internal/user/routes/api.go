package routes

import (
	c "github.com/JubaerHossain/go-service/internal/user/controllers"
	"github.com/JubaerHossain/gomd/config"
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

	v1 := Router.Group("api/v1")
	v1.GET("users", c.UserIndex())
	v1.POST("users", c.UserCreate())
	v1.GET("users/:userId", c.UserShow())
	v1.PUT("users/:userId", c.UserUpdate())
	v1.DELETE("users/:userId", c.UserDelete())

}
