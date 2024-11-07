package model

import "github.com/google/uuid"

type Vehicle struct {
	Id     uuid.UUID
	Brand  string
	Model  string
	Color  string
	Amount string
}
