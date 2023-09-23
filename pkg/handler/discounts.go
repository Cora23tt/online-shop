package handler

import (
	"net/http"
	"strconv"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getDiscounts(c *gin.Context) {
	discounts, err := h.services.Discounts.GetAll()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"discounts": discounts})
}

func (h *Handler) getDiscount(c *gin.Context) {
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
	discount, err := h.services.Discounts.GetByID(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, gin.H{"discount": discount})
}

func (h *Handler) createDiscount(c *gin.Context) {
	var input onlinedilerv3.DiscountInput
	if err := c.BindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := h.services.Discounts.Create(input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) updateDiscount(c *gin.Context) {
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
	var input onlinedilerv3.DiscountInput
	if err := c.BindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := h.services.Discounts.Update(id, input); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) deleteDiscount(c *gin.Context) {
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
	if err := h.services.Discounts.Delete(id); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
