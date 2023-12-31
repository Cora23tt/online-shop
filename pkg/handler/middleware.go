package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorisationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorisationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

// func getUserId(c *gin.Context) (int, error) {

// 	id, ok := c.Get(userCtx)
// 	if !ok {
// 		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
// 		return 0, errors.New("user id not found")
// 	}

// 	idInt, ok := id.(int)
// 	if !ok {
// 		NewErrorResponse(c, http.StatusInternalServerError, "user id is invalid type")
// 		return 0, errors.New("user id is invalid type")
// 	}

// 	return idInt, nil
// }
