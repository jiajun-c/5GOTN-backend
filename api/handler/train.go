package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"otn/services/core"
)

func Train(c *gin.Context) {
	addr := viper.GetString("rpc.addr")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	data, err := core.PutTrainData(conn)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func StartTrain(c *gin.Context) {
	addr := viper.GetString("rpc.addr")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	resp, err := core.StartTrain(conn)
	if err != nil {
		log.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}
