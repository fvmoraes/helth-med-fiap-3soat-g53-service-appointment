package domain

import "time"

type Appointment struct {
	ID          uint      `gorm:"primary_key"`
	DoctorID    uint      `gorm:"not null"`
	PatientID   uint      `gorm:"not null"`
	DateTime    time.Time `gorm:"not null"`
	Status      string    `gorm:"size:100;not null"`
	Description string    `gorm:"size:255"`
	Link        string    `gorm:"size:255"`
}
