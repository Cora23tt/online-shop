package handler

import (
	"net/http"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Email    string `json:"email" binging:"required"`
	Password string `json:"password" binging:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input onlinedilerv3.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorisation.CreateUser(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorisation.GenerateToken(input.Email, input.Password)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
