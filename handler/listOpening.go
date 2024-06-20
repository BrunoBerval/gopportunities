package handler

import (
	"net/http"

	"github.com/BrunoBerval/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Basepath /api/v1

// @Summary List openings
// @Description Show a list of all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusNotFound, "error listing openings")
		return
	}
	sendSuccess(c, "list-openings", openings)
}
