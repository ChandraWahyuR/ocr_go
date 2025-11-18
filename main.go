package main

import (
	"fmt"
	"log"
	"net"
	"parkir/app"
	"parkir/config"
	"parkir/db/migrations"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func grpcListen() <-chan error {
	out := make(chan error)
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			out <- fmt.Errorf("failed to listen gRPC: %w", err)
			return
		}

		grpcServer := grpc.NewServer()

		// pb.RegisterOcrServiceServer(grpcServer, &handler.OcrService{})

		log.Println("[gRPC] Listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			out <- fmt.Errorf("gRPC server error: %w", err)
		}
	}()
	return out
}

func restListen() <-chan error {
	out := make(chan error)
	go func() {
		router := gin.Default()
		cfg := config.EnvFile()
		db, err := config.InitDatabase(*cfg)
		if err != nil {
			out <- fmt.Errorf("gagal menghubungkan ke database: %w", err)
			return
		}

		logger := logrus.New()
		bootstrap := &app.Boostrap{
			App:    router,
			DB:     db,
			Log:    logger,
			Config: cfg,
		}
		app.App(bootstrap)

		// Inisialisasi tabel
		err = migrations.CreateTable(db)
		if err != nil {
			log.Fatal("Gagal membuat tabel:", err)
		}

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "ping"})
		})

		if err := router.Run(":8080"); err != nil {
			out <- fmt.Errorf("REST server error: %w", err)
		}
	}()
	return out
}
func main() {
	// grpcErr := grpcListen()
	restErr := restListen()

	// Tunggu salah satu error
	select {
	// case err := <-grpcErr:
	// log.Fatalf("[FATAL] gRPC crashed: %v", err)
	case err := <-restErr:
		log.Fatalf("[FATAL] REST crashed: %v", err)
	}
}
