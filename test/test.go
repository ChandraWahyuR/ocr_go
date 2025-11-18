package main

import (
	"context"
	"fmt"
	"log"
	ocrpb "parkir/internal/ocr/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gagal konek ke grpc python: %v", err)
	}
	client := ocrpb.NewFileServiceClient(conn)

	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatal("gagal mulai stream:", err)
	}

	req := &ocrpb.FileUploadRequest{
		FileName: "test.jpg",
		Chunk:    []byte("fakeimagebytes"),
	}
	if err := stream.Send(req); err != nil {
		log.Fatal("gagal kirim req:", err)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("gagal terima response:", err)
	}

	fmt.Println("Result:", res.ResultText)
}
