package main

import (
	"context"
	"k8s_demo_client/pb"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Change this for your own project

func main() {
	conn, err := grpc.Dial("gcd-service.default:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	cli := pb.NewGCDServiceClient(conn)
	req := &pb.GCDRequest{A: 2, B: 2}

	// r := gin.Default()
	// r.GET("/gcd", func(c *gin.Context) {
	// 	if res, err := cli.Compute(c, req); err == nil {
	// 		logrus.Info(res)
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"result": fmt.Sprint(res.Result),
	// 		})
	// 	} else {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	}
	// })
	// logrus.Info("start client")
	// if err := r.Run(":3000"); err != nil {
	// 	log.Fatalf("Failed to run server: %v", err)
	// }

	for {
		time.Sleep(1 * time.Second)
		res, err := cli.Compute(context.Background(), req)
		logrus.Infof("Compute: res=%v, err=%v", res, err)
	}
}
