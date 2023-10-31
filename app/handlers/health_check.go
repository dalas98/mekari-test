package handlers

import (
	"net/http"

	"github.com/dalas98/mekari-test/helpers"
	"github.com/gin-gonic/gin"
)

func (h *APIGOHandler) Check(c *gin.Context) {
	hc, err := h.APIGOUsecase.HealthCheck()
	if err != nil {
		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(c, http.StatusOK, hc, nil)
}
