package domain

type AppointmentRepository interface {
	Create(appointment *Appointment) error
	FindByID(id uint) (*Appointment, error)
	Update(appointment *Appointment) error
	Delete(id uint) error
}
