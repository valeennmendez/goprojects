package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valeennmendez/api-go/connection"
	"github.com/valeennmendez/api-go/models"
)

func TotalPatientsData(c *gin.Context) {
	var count int64

	connection.DB.Model(&models.Patients{}).Count(&count)
	log.Println("GetTotalPatients llamada")
	c.JSON(http.StatusAccepted, gin.H{
		"total": count,
	})
}

func AppointmentToday(c *gin.Context) {
	var count int64

	today := time.Now().Format("2006-01-02")

	err := connection.DB.Raw("SELECT COUNT(*) FROM appoinments WHERE DATE(fecha) = ?", today).Scan(&count).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "The date no exist.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})

}
