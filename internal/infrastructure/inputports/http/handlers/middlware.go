package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

type UserIdentity interface {
	UserIdentity(c *gin.Context)
}

func getID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return 0, errors.New("id not found")
	}

	return id, nil
}

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	claims, err := h.service.UserService.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("id", claims.UserID)
	c.Set("role", claims.Role)
}

func isUserAuthorized(c *gin.Context, inputID, id int) bool {
	role, ok := c.Get("role")
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, roleDoesNotExist)
		return false
	}

	if inputID != id && role.(int) <= User {
		return false
	}

	return true
}
