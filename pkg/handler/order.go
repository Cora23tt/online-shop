package handler

import (
	"net/http"
	"strconv"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getOrders(c *gin.Context) {
	orders, err := h.services.Orders.GetAll()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"discounts": orders})
}

func (h *Handler) getOrder(c *gin.Context) {
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
	order, err := h.services.Orders.GetByID(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, gin.H{"discount": order})
}

func (h *Handler) createOrder(c *gin.Context) {
	var input onlinedilerv3.Order
	if err := c.BindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := h.services.Orders.Create(input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) updateOrder(c *gin.Context) {
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
	var input onlinedilerv3.Order
	if err := c.BindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := h.services.Orders.Update(id, input); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) deleteOrder(c *gin.Context) {
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
	if err := h.services.Orders.Delete(id); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
