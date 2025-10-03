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

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `json:"username"`
	PasswordHash string         `json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// Authentication request/response models
type LoginRequest struct {
	Username string `json:"username" example:"yunus"`
	Password string `json:"password" example:"yunus"`
}

type LoginResponse struct {
	Token     string    `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ExpiresAt time.Time `json:"expires_at" example:"2025-10-03T20:30:00Z"`
}

// Success response models
type CreateFaultReportSuccess struct {
	Message string     `json:"message" example:"Fault report created successfully"`
	Report  ReportData `json:"report"`
}

type ReportData struct {
	ID              uint      `json:"id" example:"1"`
	Title           string    `json:"title" example:"Machine malfunction"`
	Asset           string    `json:"asset" example:"Conveyor Belt A1"`
	MachineID       string    `json:"machine_id" example:"MCH001"`
	Priority        string    `json:"priority" example:"High"`
	UserDescription string    `json:"user_description" example:"Machine stopped working"`
	Timestamp       time.Time `json:"timestamp" example:"2025-10-03T08:30:00Z"`
	ReportedBy      string    `json:"reported_by" example:"john_doe"`
}

type GetFaultReportsSuccess struct {
	Message string        `json:"message" example:"Fault reports retrieved successfully"`
	Data    []FaultReport `json:"data"`
	Count   int           `json:"count" example:"5"`
}

type GetFaultReportSuccess struct {
	Message string      `json:"message" example:"Fault report retrieved successfully"`
	Data    FaultReport `json:"data"`
}

// Error response models
type ErrorResponse struct {
	Error string `json:"error"`
}

type InvalidRequestError struct {
	Error string `json:"error" example:"Invalid request"`
}

type InvalidCredentialsError struct {
	Error string `json:"error" example:"Invalid credentials"`
}

type TokenGenerationError struct {
	Error string `json:"error" example:"Could not generate token"`
}

type ParseJSONError struct {
	Error string `json:"error" example:"Cannot parse JSON"`
}

type MissingFieldsError struct {
	Error string `json:"error" example:"Missing required fields"`
}

type UnauthorizedError struct {
	Error string `json:"error" example:"Unauthorized"`
}

type CreateReportError struct {
	Error string `json:"error" example:"Could not create fault report"`
}

type RetrieveReportsError struct {
	Error string `json:"error" example:"Could not retrieve fault reports"`
}

type ReportNotFoundError struct {
	Error string `json:"error" example:"Fault report not found"`
}
