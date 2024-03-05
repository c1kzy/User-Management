package handlers

import (
	"net/http"
	"restapi/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

// @Summary UpdateUser
// @Security ApiKeyAuth
// @Tags users
// @Description update user
// @Accept  json
// @Produce  json
// @Param input body domain.User true "user info"
// @Success 200 {object} errResponse
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/users/:id [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var input domain.User

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

	if !isUserAuthorized(c, input.ID, id) {
		NewErrorResponse(c, http.StatusInternalServerError, permissionDenied)
		return
	}

	err = h.service.UserService.Update(id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"errResponse": err,
	})
}

// @Summary DeleteUser
// @Security ApiKeyAuth
// @Tags users
// @Description delete user
// @Accept  json
// @Produce  json
// @Success 200 {object} errResponse
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/users/:id [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	var input domain.User

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

	if !isUserAuthorized(c, input.ID, id) {
		NewErrorResponse(c, http.StatusInternalServerError, permissionDenied)
		return
	}

	err = h.service.UserService.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"errResponse": err,
	})
}
