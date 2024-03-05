package handlers

import (
	"net/http"
	"restapi/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

// @Summary CreatePost
// @Security ApiKeyAuth
// @Tags posts
// @Description create post
// @Accept  json
// @Produce  json
// @Param input body domain.Post true "post info"
// @Success 200 {object} errResponse
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/users/posts/ [post]
func (h *Handler) CreatePost(c *gin.Context) {
	var input domain.Post

	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, ok := c.Get("id")
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, roleDoesNotExist)
		return
	}

	if !isUserAuthorized(c, input.UserID, id.(int)) {
		NewErrorResponse(c, http.StatusInternalServerError, permissionDenied)
		return
	}

	err := h.service.PostService.CreatePost(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"error": err,
	})
}

// @Summary UpdatePost
// @Security ApiKeyAuth
// @Tags posts
// @Description update post
// @Accept  json
// @Produce  json
// @Param input body domain.Post true "post info"
// @Success 200 {object} errResponse
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/users/posts/:id [put]
func (h *Handler) UpdatePost(c *gin.Context) {
	var input domain.Post

	id, err := getID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, ok := c.Get("id")
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, roleDoesNotExist)
		return
	}

	if !isUserAuthorized(c, input.UserID, userID.(int)) {
		NewErrorResponse(c, http.StatusInternalServerError, permissionDenied)
		return
	}

	err = h.service.PostService.UpdatePost(id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"errResponse": err,
	})
}

// @Summary DeletePost
// @Security ApiKeyAuth
// @Tags posts
// @Description delete post
// @Accept  json
// @Produce  json
// @Param input body domain.Post true "post info"
// @Success 200 {object} errResponse
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/users/posts/:id [delete]
func (h *Handler) DeletePost(c *gin.Context) {
	var input domain.Post

	id, err := getID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, ok := c.Get("id")
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, roleDoesNotExist)
		return
	}

	if !isUserAuthorized(c, input.UserID, userID.(int)) {
		NewErrorResponse(c, http.StatusInternalServerError, permissionDenied)
		return
	}

	err = h.service.PostService.DeletePost(id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"errResponse": err,
	})

}

// @Summary ListPosts
// @Security ApiKeyAuth
// @Tags posts
// @Description get public posts
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.PublicPost
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/users/posts/:page [get]
func (h *Handler) ListPosts(c *gin.Context) {
	page := c.Param("page")

	posts, err := h.service.ListPosts(page)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"posts": posts,
	})
}
