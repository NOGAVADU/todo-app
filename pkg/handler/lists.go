package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nogavadu/todo-app"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Create List
// @Security ApiKeyAuth
// @Tags lists
// @Description create new list
// @ID create-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} todo.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:listId [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	list, err := h.services.TodoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update List
// @Security ApiKeyAuth
// @Tags lists
// @Description update list
// @ID update-list
// @Accept  json
// @Produce  json
// @Success 200 {object} StatusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:listId [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	var input todo.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Update(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// @Summary Delete List
// @Security ApiKeyAuth
// @Tags lists
// @Description delete list
// @ID delete-list
// @Accept  json
// @Produce  json
// @Success 200 {object} StatusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:listId [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	err = h.services.TodoList.Delete(userId, listId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
