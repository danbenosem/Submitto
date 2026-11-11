package models

import "time"

type Submission struct {
	ID int
	FormId int
	SubmittedAt time.Time
}

