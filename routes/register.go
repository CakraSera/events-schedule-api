package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// RegisterForEvent godoc
// @Summary Register for an event
// @Description Register the authenticated user for a specific event
// @Tags registrations
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id}/register [post]
func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed register to database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

// CancelRegistration godoc
// @Summary Cancel event registration
// @Description Cancel the authenticated user's registration for a specific event
// @Tags registrations
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id}/register [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegister(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could no cancel registeration on database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success delete data"})
}
