package handler

import (
	"net/http"
	"strconv"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllConsignments(c *gin.Context) {
	consignments, err := h.services.Consignments.GetAll()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"consignments": consignments})
}

func (h *Handler) getConsignment(c *gin.Context) {
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
	consignment, err := h.services.Consignments.GetByID(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, gin.H{"consignment": consignment})
}

func (h *Handler) createConsignment(c *gin.Context) {
	var input onlinedilerv3.Consignment
	if err := c.BindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := h.services.Consignments.Create(input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) updateConsignment(c *gin.Context) {
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
	var input onlinedilerv3.Consignment
	if err := c.BindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err = h.services.Consignments.Update(id, input); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) deleteConsignment(c *gin.Context) {
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
	if err := h.services.Consignments.Delete(id); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}
