package model

import "time"

type Event struct {
	ID          int
	EventName   string
	Description string
	CreateDate  time.Time
	StartDate   time.Time
	EndDate     time.Time
	Status      bool
	Type        bool
	Image       string
}
