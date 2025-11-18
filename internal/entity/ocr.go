package entity

import "time"

type Ocr struct {
	ID        string
	PlatNomor string
	IsExit    bool
	Entered   time.Time
	Exited    time.Time
}
