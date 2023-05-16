package core

import (
	"context"
	"google.golang.org/grpc"
	"otn/pb"
	"time"
)

func Alert(conn *grpc.ClientConn, req *pb.AlertRequest) (*pb.AlertResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	c := pb.NewCoreClient(conn)
	alert, err := c.Alert(ctx, req)
	if err != nil {
		return nil, err
	}
	return alert, nil
}
