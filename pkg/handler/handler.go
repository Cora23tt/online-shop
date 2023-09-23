package handler

import (
	"github.com/cora23tt/onlinedilerv3/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api/:lang", h.userIdentity)
	{
		product := api.Group("/product")
		{
			product.GET("/", h.getProducts)
			product.GET("/:id", h.getProduct)
			product.POST("/", h.createProduct)
			product.PATCH("/:id", h.updateProduct)
			product.DELETE("/:id", h.deleteProduct)
		}
		category := api.Group("/category")
		{
			category.GET("/", h.getAllCategories)
			category.GET("/:id", h.getCategory)
			category.POST("/", h.createCategory)
			category.DELETE("/:id", h.deleteCategory)
			category.PATCH("/:id", h.updateCategory)
		}
		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUser)
			users.DELETE("/:id", h.deleteUser)
			users.PATCH("/:id", h.updateUser)
		}
		consignments := api.Group("/consignments")
		{
			consignments.GET("/", h.getAllConsignments)
			consignments.GET("/:id", h.getConsignment)
			consignments.POST("/", h.createConsignment)
			consignments.PATCH("/:id", h.updateConsignment)
			consignments.DELETE("/:id", h.deleteConsignment)
		}
		discounts := api.Group("/discounts")
		{
			discounts.GET("/", h.getDiscounts)
			discounts.GET("/:id", h.getDiscount)
			discounts.POST("/", h.createDiscount)
			discounts.PATCH("/:id", h.updateDiscount)
			discounts.DELETE("/:id", h.deleteDiscount)
			forClients := discounts.Group("/for-clients")
			{
				forClients.GET("/", h.getClientsDiscounts)
				forClients.GET("/:id", h.getClientDiscount)
				forClients.POST("/", h.createClientDiscount)
				forClients.PATCH("/:id", h.updateClientDiscount)
				forClients.DELETE("/:id", h.deleteClientDiscount)
			}
		}
		orders := api.Group("/orders")
		{
			orders.GET("/", h.getOrders)
			orders.GET("/:id", h.getOrder)
			orders.POST("/", h.createOrder)
			orders.PATCH("/:id", h.updateOrder)
			orders.DELETE("/:id", h.deleteOrder)
			items := orders.Group("/items/:order_id")
			{
				items.GET("/", h.getItems)
				items.POST("/", h.addItem)
				items.PATCH("/", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}

	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS,GET,PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
