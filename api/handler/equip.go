package handler

import (
	"github.com/gin-gonic/gin"
	"otn/dal"
	"otn/services/crud"
	"strconv"
)

func InsertEquip(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	err := crud.InsertEquip(&dal.Equip{
		Id:   int64(id),
		Name: name,
	})
	if err != nil {
		return
	}
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func DeleteEquip(ctx *gin.Context) {
	sid := ctx.Query("id")
	id, _ := strconv.Atoi(sid)
	err := crud.DeleteEquip(int64(id))
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func AllEquip(ctx *gin.Context) {
	equips, err := crud.GetAllEquip()
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, equips)
}

func OneEquip(ctx *gin.Context) {
	sid := ctx.Query("id")
	id, _ := strconv.Atoi(sid)
	e, err := crud.GetOneEquip(int64(id))
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, e)
}

func UpdateEquip(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	err := crud.UpdateEquip(&dal.Equip{
		Id:   int64(id),
		Name: name,
	})
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
