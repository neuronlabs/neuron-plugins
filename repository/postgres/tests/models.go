package tests

import (
	"time"
)

//go:generate neuron-generator models methods --format=gofmt --single-file .

// InsertModel is the model prepared for insertion.
type SimpleModel struct {
	ID        int        `neuron:"type=primary"`
	Attr      string     `neuron:"type=attr"`
	CreatedAt *time.Time `neuron:"type=attr"`
}

// OmitModel is the model with omitted field.
type OmitModel struct {
	ID        int
	OmitField string `db:"-"`
}

// Model is the generic model used for insertion.
type Model struct {
	ID int `neuron:"type=primary"`

	AttrString string  `neuron:"type=attr"`
	StringPtr  *string `neuron:"type=attr"`

	Int       int        `neuron:"type=attr"`
	CreatedAt time.Time  `neuron:"type=attr"`
	UpdatedAt *time.Time `neuron:"type=attr"`
	DeletedAt *time.Time `neuron:"type=attr"`
}
