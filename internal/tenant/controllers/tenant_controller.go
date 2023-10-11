package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/JubaerHossain/gomd/gomd"
	"github.com/JubaerHossain/go-service/internal/tenant/validation"
	"github.com/JubaerHossain/go-service/internal/tenant/services"
	"net/http"
)
func TenantIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
        page := c.DefaultQuery("page", "1")
        limit := c.DefaultQuery("limit", "10")
        status := c.DefaultQuery("status", "")

        var filter map[string]interface{} = make(map[string]interface{})
        filter["page"] = page
        filter["limit"] = limit
        filter["status"] = status

        tenants, paginate := services.AllTenant(filter)

        gomd.Res.Code(200).Data(tenants).Raw(map[string]interface{}{
            "meta": paginate,
        }).Json(c)
	}
}


func TenantCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createTenant validation.CreateTenantRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createTenant); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		tenant := services.CreateATenant(createTenant)

		gomd.Res.Code(http.StatusCreated).Message("success").Data(tenant).Json(c)
	}
}


func TenantShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		tenantId := c.Param("tenantId")

		tenant := services.ATenant(tenantId)

		gomd.Res.Code(http.StatusOK).Message("success").Data(tenant).Json(c)
	}
}


func TenantUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateTenant validation.UpdateTenantRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		tenantId := c.Param("tenantId")

		if err := c.ShouldBind(&updateTenant); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		tenant, err := services.UpdateATenant(tenantId, updateTenant)

		if err != nil {
			gomd.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(tenant).Json(c)
	}
}


func TenantDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		tenantId := c.Param("tenantId")
		err := services.DeleteATenant(tenantId)

		if !err {
			gomd.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}