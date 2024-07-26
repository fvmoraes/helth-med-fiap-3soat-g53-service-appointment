package repository

import (
	"helthmed-appointment/internal/appointment/domain"
	"helthmed-appointment/internal/appointment/infrastructure/db"
)

type AppointmentRepository struct{}

func NewAppointmentRepository() domain.AppointmentRepository {
	return &AppointmentRepository{}
}

func (r *AppointmentRepository) Create(appointment *domain.Appointment) error {
	return db.DB.Create(appointment).Error
}

func (r *AppointmentRepository) FindByID(id uint) (*domain.Appointment, error) {
	var appointment domain.Appointment
	result := db.DB.First(&appointment, id)
	return &appointment, result.Error
}

func (r *AppointmentRepository) Update(appointment *domain.Appointment) error {
	return db.DB.Save(appointment).Error
}

func (r *AppointmentRepository) Delete(id uint) error {
	return db.DB.Delete(&domain.Appointment{}, id).Error
}
