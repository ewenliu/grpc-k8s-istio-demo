package main

import (
	"context"
	"k8s_demo_client/pb"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// 通过 k8s 的 endpoint 进行服务发现
	conn, err := grpc.Dial("gcd-service:3000", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("Dial failed: %v", err)
	}

	cli := pb.NewGCDServiceClient(conn)
	req := &pb.GCDRequest{A: 2, B: 2}

	// 持续 rpc 调用远端的 Compute 函数
	for {
		time.Sleep(1 * time.Second)
		res, err := cli.Compute(context.Background(), req)
		logrus.Infof("Compute: res=%v, err=%v", res, err)
	}
}
