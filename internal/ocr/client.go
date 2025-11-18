package ocr

import (
	"context"
	"io"
	"mime/multipart"
	ocrpb "parkir/internal/ocr/pb"
	"strconv"
)

type ocrProcessor struct {
	client ocrpb.FileServiceClient
}

func NewOCRProcessor(client ocrpb.FileServiceClient) *ocrProcessor {
	return &ocrProcessor{client: client}
}

func (p *ocrProcessor) UploadFile(ctx context.Context, file multipart.File, fileName string) (string, float64, error) {
	stream, err := p.client.Upload(ctx) // kita buka pintu stream, disini kita client karena memerintah py suruh proses datanya
	if err != nil {
		return "", 0, err
	}

	defer file.Close()

	buf := make([]byte, 1024*32) // Membuat buffer untuk baca potongan file sebesar 32 KB per loop.
	// kan data dikirim data, jadi misal data 7mb, nanti 7000kb/32 jadi ada 218 loop
	for { // loop yang di jelaskan tadi
		n, err := file.Read(buf)
		if err == io.EOF { // jika selesai baca file break
			break
		}
		if err != nil {
			return "", 0, err
		}

		req := &ocrpb.FileUploadRequest{
			FileName: fileName,
			Chunk:    buf[:n],
		}

		if err := stream.Send(req); err != nil { // stream.Send(req) adalah gRPC call per chunk.
			return "", 0, err
		}
	}

	res, err := stream.CloseAndRecv() // Setelah semua chunk dikirim, tutup stream, lalu tunggu response dari server. Response ini hanya muncul sekali, dan berisi hasil OCR:
	if err != nil {
		return "", 0, err
	}

	conf, _ := strconv.ParseFloat(res.Accuracy, 64) // Respone py ada text dan accuracy. accuracy dijadikan float setelah di proses di protobuf.

	return res.ResultText, conf, nil
}

// Setelah itu ke usecase agar dihubungkan
// jangan lupa sisi python sebagai server buat juga grcpnya
