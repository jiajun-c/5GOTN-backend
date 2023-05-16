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
	//flag.Parse()
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//defer conn.Close()
	//
	//c := pb.NewCoreClient(conn)
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	//defer cancel()
	//
	//r, err := c.Train(ctx, &pb.TrainRequest{Active: true})
	//if err != nil {
	//	log.Error(err)
	//}
	//fmt.Println(r.Success)
	//
	//data1 := pb.DataStruct{
	//	Clogid:        0,
	//	Calarmcode:    0,
	//	Cneid:         0,
	//	Calarmlevel:   0,
	//	Coccurutctime: "aa",
	//	Clocationinfo: "bb",
	//	Clineport:     "cc",
	//}
	//var datas = []*pb.DataStruct{&data1}
	//var dataset = pb.AnalyseRequest{Data: datas}
	//r0, err := c.Analyse(ctx, &dataset)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//println(r0)
	////var local1 = &pb.Datainput{Clocationinfo: "aann"}
	//var local1 = &pb.Datainput{Clocationinfo: "1-1-aa:bb"}
	//var names = []*pb.Datainput{local1}
	//r1, err := c.Alert(ctx, &pb.AlertRequest{Datarequest: names})
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//for _, v := range r1.Dataresponse {
	//	fmt.Println(v)
	//}
	////println(r1.Dataresponse)
	r := gin.Default()
	r.Use(middlewares.Cors())

	router.Route(r)
	addr := viper.GetString("api.addr")
	r.Run(addr)
}
