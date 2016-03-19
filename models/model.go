package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type Model struct {
	ID        string     `gorm:"column:id; primary_key" sql:"type:char(255);not null;unique;primary key"`
	CreatedAt time.Time  `json:"created_at,omitempty" sql:"index"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" sql:"index"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}

func (m *Model) BeforeCreate() error {
	if m.ID == "" {
		guid := uuid.NewV4()
		m.ID = guid.String()
	}
	return nil
}
