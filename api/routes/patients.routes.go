package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valeennmendez/api-go/connection"
	"github.com/valeennmendez/api-go/models"
)

func CreatePatient(c *gin.Context) {

	var patient models.Patients

	// Decodificar el cuerpo de la solicitud en la estructura patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format: " + err.Error(),
		})
		return
	}

	// Crear el paciente en la base de datos
	if err := connection.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create patient: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Patients created succesfully",
	})

}

func EditPatient(c *gin.Context) {
	var patient models.Patients

	id := c.Param("id")

	if err := connection.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Patient not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON formart" + err.Error(),
		})
		return
	}

	if err := connection.DB.Save(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to actualice patient.",
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Patient edited succesfully",
	})

}

func GetAllPatients(c *gin.Context) {
	var patients []models.Patients

	connection.DB.Find(&patients)

	c.JSON(http.StatusAccepted, &patients)

}

func GetPatientByID(c *gin.Context) {
	var patient models.Patients

	id := c.Param("id")

	if err := connection.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Patient not found",
		})
	}

	c.JSON(http.StatusAccepted, &patient)
}

func DeletePacients(c *gin.Context) {
	var patient models.Patients

	id := c.Param("id")

	if err := connection.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Pacient not found" + err.Error(),
		})
		return
	}

	if err := connection.DB.Unscoped().Delete(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete pacient" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Patient deleted successfully",
	})
}

func SearchPatient(c *gin.Context) {
	search := c.Query("p")

	var patients []models.Patients

	if search == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "search term is required",
		})
		return
	}

	if err := connection.DB.Where("full_name LIKE ? OR dni LIKE ?", "%"+search+"%", "%"+search+"%").Find(&patients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error searching the patient" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, patients)

}
