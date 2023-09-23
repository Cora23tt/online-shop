package handler

import (
	"net/http"
	"strconv"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getProducts(c *gin.Context) {
	lang := c.Param("lang")
	queryType := c.Query("type")
	productName := c.Query("name")
	limit_s := c.Query("limit")
	offset_s := c.Query("offset")
	categoryID := c.Query("category_id")

	if lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid param lang"})
		return
	}

	var limit, offset int

	if offset_s != "" {
		offset_int, err := strconv.Atoi(offset_s)
		offset = offset_int
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid offset value"})
			return
		}
	}
	if limit_s != "" {
		limit_int, err := strconv.Atoi(limit_s)
		limit = limit_int
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid limit value"})
			return
		}
	}

	switch queryType {
	case "by_category":
		if categoryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid category name"})
			return
		}
		categoryIDint, err := strconv.Atoi(categoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid category id value"})
			return
		}
		products, err := h.services.Products.ByCategory(lang, categoryIDint)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"products": products})

	case "search":
		if productName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid product name"})
			return
		}
		if limit != 0 {
			products, err := h.services.Products.SearchWithLimit(limit, offset, lang, productName)
			if err != nil {
				c.Status(http.StatusNotFound)
				return
			}
			c.JSON(http.StatusOK, gin.H{"products": products})
			return
		}
		products, err := h.services.Products.Search(lang, productName)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, gin.H{"products": products})

	case "top_rated":
		if limit != 0 {
			products, err := h.services.Products.TopRatedWithLimit(limit, offset, lang)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
			c.JSON(http.StatusOK, gin.H{"products": products})
			return
		}
		products, err := h.services.Products.TopRated(lang)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"products": products})

	default:
		products, err := h.services.Products.GetAll(lang)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"products": products})
	}
}

func (h *Handler) createProduct(c *gin.Context) {
	var input struct {
		Product      onlinedilerv3.Product              `json:"product"`
		Translations []onlinedilerv3.ProductTranslation `json:"translations"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productID, err := h.services.Products.Create(input.Product, input.Translations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product_id": productID})
}

func (h *Handler) getProduct(c *gin.Context) {
	id_s := c.Param("id")
	lang := c.Param("lang")
	if id_s == "" || lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id or lang"})
		return
	}

	id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id number"})
		return
	}
	product, err := h.services.Products.GetByID(lang, id)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (h *Handler) updateProduct(c *gin.Context) {
	id_s := c.Param("id")
	if id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id number"})
		return
	}
	id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id number"})
		return
	}

	var input onlinedilerv3.ProductComplect
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.services.Products.Update(id, input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) deleteProduct(c *gin.Context) {
	id_s := c.Param("id")
	if id_s == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id number"})
		return
	}
	id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id number"})
		return
	}
	err = h.services.Products.Delete(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
