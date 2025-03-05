package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:listId", h.getListById)
			lists.PUT("/:listId", h.updateList)
			lists.DELETE("/:listId", h.deleteList)
		}

		items := api.Group(":listId/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:itemId", h.getItemById)
			items.PUT("/:itemId", h.updateItem)
			items.DELETE("/:itemId", h.deleteItem)
		}
	}

	return router
}
