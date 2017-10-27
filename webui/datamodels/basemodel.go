package datamodels

import (
	"time"
	"github.com/go-pg/pg/orm"
)

type BaseModel struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (m *BaseModel) BeforeInsert(db orm.DB) error {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

func (m *BaseModel) BeforeUpdate(db orm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

