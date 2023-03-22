package router

import (
	"github.com/gin-gonic/gin"
	"otn/api/handler"
)

func EquipRoute(engine *gin.Engine) {
	r := engine.Group("/equip")
	r.POST("/delete", handler.DeleteEquip)
	r.POST("/insert", handler.InsertEquip)
	r.GET("/all", handler.AllEquip)
	r.GET("/one", handler.OneEquip)
	r.POST("/update", handler.UpdateEquip)
}
