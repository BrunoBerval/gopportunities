package handler

import (
	"net/http"

	"github.com/BrunoBerval/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusNotFound, "error listing openings")
		return
	}
	sendSuccess(c, "list-openings", openings)
}
