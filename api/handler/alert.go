package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"otn/pb"
	"otn/services/core"
)

func Alert(ctx *gin.Context) {
	addr := viper.GetString("rpc.addr")
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	locate := ctx.PostForm("locate")
	ans, err := core.Alert(conn, &pb.AlertRequest{Datarequest: []*pb.Datainput{&pb.Datainput{Clocationinfo: locate}}})
	if err != nil {
		ctx.JSON(200, gin.H{"error": err})
		return
	}
	ctx.JSON(200, ans)
}
