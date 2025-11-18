package model

import "time"

type Ocr struct {
	ID        string    `json:"id"`
	PlatNomor string    `json:"plat_nomor"`
	IsExit    bool      `json:"is_exit"`
	Entered   time.Time `json:"entered"`
	Exited    time.Time `json:"exited"`
}

type OcrPy struct {
	Text       string
	Confidence float64
}
