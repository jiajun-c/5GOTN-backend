package router

import (
	"github.com/gin-gonic/gin"
	"otn/api/handler"
	"otn/middlewares"
)

func Route(engine *gin.Engine) {
	auth := middlewares.Auth()
	engine.POST("/login", auth.LoginHandler)
	admin := engine.Group("/admin")
	admin.POST("/refresh_token", auth.RefreshHandler)
	admin.POST("/check_token", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})
	admin.Use(auth.MiddlewareFunc())
	equip := admin.Group("/equip")
	equip.POST("/delete", handler.DeleteEquip)
	equip.POST("/insert", handler.InsertEquip)
	equip.GET("/all", handler.AllEquip)
	equip.GET("/one", handler.OneEquip)
	equip.POST("/update", handler.UpdateEquip)

	record := admin.Group("/record")
	record.DELETE("/delete", handler.DeleteRecord)
	record.POST("/insert", handler.InsertRecord)
	record.GET("/all", handler.GetAllRecords)
	record.GET("/some", handler.FilterRecord)
	AdminRoute(engine)
}
