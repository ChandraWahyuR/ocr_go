package config

import (
	"log"
	ocrpb "parkir/internal/ocr/pb"

	"google.golang.org/grpc"
)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}

	return ocrpb.NewFileServiceClient(conn)
}
