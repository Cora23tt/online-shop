package handler

import (
	"net/http"
	"strconv"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	queryType := c.Query("type")
	name := c.Query("name")

	switch queryType {
	case "search":
		if name == "" {
			c.Status(http.StatusBadRequest)
			return
		}
		users, err := h.services.Users.Search(name)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"users": users})

	default:
		users, err := h.services.Users.GetAll()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func (h *Handler) getUser(c *gin.Context) {
	id_s := c.Param("id")
	if id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	User, err := h.services.Users.GetByID(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, gin.H{"user": User})
}

func (h *Handler) deleteUser(c *gin.Context) {
	id_s := c.Param("id")
	if id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	err = h.services.Users.Delete(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *Handler) updateUser(c *gin.Context) {
	id_s := c.Param("id")
	if id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var input onlinedilerv3.User
	err = c.BindJSON(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user input"})
		return
	}
	err = h.services.Users.Update(id, input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated successfuly"})
}
