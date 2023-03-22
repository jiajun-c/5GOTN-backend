package main

import (
	"otn/api/router"
	"otn/config"
	"otn/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.LoadDB()
}

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	router.Route(r)
	addr := viper.GetString("api.addr")
	r.Run(addr)
}
