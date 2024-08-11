package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Appoinment struct {
	gorm.Model
	PacienteID     int            `json:"pacienteid" gorm:"foreignkey:PacienteID"`
	Fecha          datatypes.Date `json: "fecha" gorm:"type:DATE"`
	Hora           string         `json: "hora"`
	MotivoConsulta string         `json: "motivoconsulta"`
	Paciente       Patients       `json: "paciente" gorm:"association_foreignkey:PacienteID"`
}
