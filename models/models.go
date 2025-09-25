package models

import (
	"time"

	"gorm.io/gorm"
)

type FaultReport struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Priority        string         `json:"priority"`
	Asset           string         `json:"asset"`
	Title           string         `json:"title"`
	Timestamp       time.Time      `json:"timestamp"`
	Duration        string         `json:"duration"`
	UserDescription string         `json:"user_description"`
	ReportedBy      string         `json:"reported_by"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
