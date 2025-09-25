package entity

import "time"

type Encounter struct {
	UserID int64
	PartnerID int64
	Timestamp time.Time
}