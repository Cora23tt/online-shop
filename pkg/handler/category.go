package handler

import (
	"net/http"
	"strconv"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllCategories(c *gin.Context) {
	lang := c.Param("lang")
	queryType := c.Query("type")
	categoryName := c.Query("name")

	switch queryType {

	case "search":
		if categoryName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid product name"})
			return
		}
		categoies, err := h.services.Categories.Search(lang, categoryName)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusFound, gin.H{"products": categoies})

	default:
		categories, err := h.services.Categories.GetAll(lang)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"categories": categories})
	}
}

func (h *Handler) getCategory(c *gin.Context) {
	lang := c.Param("lang")
	id := c.Param("id")

	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	category, err := h.services.Categories.Get(lang, id_int)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": category})
}

func (h *Handler) createCategory(c *gin.Context) {
	var input onlinedilerv3.Category
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, err := h.services.Categories.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) deleteCategory(c *gin.Context) {
	id_param := c.Param("id")
	if id_param == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(id_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	err = h.services.Categories.Delete(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) updateCategory(c *gin.Context) {
	id_param := c.Param("id")
	if id_param == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(id_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var input onlinedilerv3.Category
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.services.Categories.Update(id, input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
