package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"otn/api/router"
	"otn/config"
	"otn/middlewares"
)

func init() {
	config.LoadDB()
}

var (
	addr = flag.String("addr", "127.0.0.1:8000", "the address to connect to")
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())

	router.Route(r)
	addr := viper.GetString("api.addr")
	r.Run(addr)
}
