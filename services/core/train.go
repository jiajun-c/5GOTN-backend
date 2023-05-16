package core

import (
	"context"
	"google.golang.org/grpc"
	"otn/pb"
	"otn/services/crud"
	"time"
)

type TrainResp struct {
	allCount int
	level0   int
	level1   int
	level2   int
	level3   int
}

func PutTrainData(conn *grpc.ClientConn) (*pb.AnalyseResponse, error) {
	recordSets, _ := crud.GetAllRecords()
	var trainSets []*pb.DataStruct
	for _, record := range recordSets {
		trainSets = append(trainSets, &pb.DataStruct{
			Clogid:        int32(record.DeviceID),
			Calarmcode:    int32(record.WarningCode),
			Cneid:         int32(record.DeviceID),
			Calarmlevel:   int32(record.WarningLevel),
			Coccurutctime: record.WarningTime.Format("2006/01/02 15:04"),
			Clocationinfo: record.DeviceAddress,
			Clineport:     record.DevicePort,
		})
	}
	c := pb.NewCoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	analyse, err := c.Analyse(ctx, &pb.AnalyseRequest{Data: trainSets})
	if err != nil {
		return nil, err
	}

	return analyse, err
}

func StartTrain(conn *grpc.ClientConn) (*pb.TrainResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	c := pb.NewCoreClient(conn)

	trainResp, err := c.Train(ctx, &pb.TrainRequest{Active: true})
	if err != nil {
		return nil, err
	}
	return trainResp, err
}
