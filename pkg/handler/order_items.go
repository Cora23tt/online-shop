package handler

import (
	"net/http"
	"strconv"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getItems(c *gin.Context) {
	id_s := c.Param("order_id")
	if id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	order_id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	items, err := h.services.OrderItems.GetItems(order_id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"discounts": items})
}

func (h *Handler) addItem(c *gin.Context) {
	id_s := c.Param("order_id")
	if id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	order_id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var input onlinedilerv3.OrderItem
	if err := c.BindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	items, err := h.services.OrderItems.Add(order_id, input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) updateItem(c *gin.Context) {
	id_s := c.Param("order_id")
	if id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	order_id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	var input onlinedilerv3.OrderItem
	if err := c.BindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := h.services.OrderItems.Update(order_id, input); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) deleteItem(c *gin.Context) {
	order_id_s := c.Param("order_id")
	item_id_s := c.Param("order_id")
	if order_id_s == "" || item_id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	order_id, err := strconv.Atoi(order_id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	item_id, err := strconv.Atoi(item_id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	if err := h.services.OrderItems.Delete(order_id, item_id); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
