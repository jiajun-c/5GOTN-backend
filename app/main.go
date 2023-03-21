package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"otn/api/router"
	"otn/config"
)

func init() {
	config.LoadDB()
}

func main() {
	r := gin.Default()
	router.Route(r)
	addr := viper.GetString("api.addr")
	r.Run(addr)
}
