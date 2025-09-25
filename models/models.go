package models

import (
	"time"

	"gorm.io/gorm"
)

type Problems struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Priority  string         `json:"priority"`
	Asset     string         `json:"asset"`
	Title     string         `json:"title"`
	Timestamp time.Time      `json:"timestamp"`
	Duration  string         `json:"duration"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type FaultReport struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Priority        string         `json:"priority"`
	Asset           string         `json:"asset"`
	Title           string         `json:"title"`
	MachineID       string         `json:"machine_id"`
	Timestamp       time.Time      `json:"timestamp"`
	UserDescription string         `json:"user_description"`
	ReportedBy      string         `json:"reported_by"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// Gelen istek için kullanılacak yapı
type CreateFaultReportRequest struct {
	Title           string `json:"title"`
	UserDescription string `json:"user_description"`
	Clock           int64  `json:"clock"`
	MachineID       string `json:"machine_id"`
	Asset           string `json:"asset"`
}

// API yanıtı için kullanılacak yapı
type FaultReportResponse struct {
	Priority        string    `json:"priority"`
	Asset           string    `json:"asset"`
	Title           string    `json:"title"`
	Duration        string    `json:"duration"`
	Timestamp       time.Time `json:"timestamp"`
	UserDescription string    `json:"user_description"`
}
