package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nogavadu/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewService(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:listId", h.getListById)
			lists.PUT("/:listId", h.updateList)
			lists.DELETE("/:listId", h.deleteList)

			items := lists.Group("/:listId/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:itemId", h.getItemById)
			items.PUT("/:itemId", h.updateItem)
			items.DELETE("/:itemId", h.deleteItem)
		}
	}

	return router
}
