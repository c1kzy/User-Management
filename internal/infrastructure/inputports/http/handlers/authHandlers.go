package handlers

import (
	"net/http"
	"restapi/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body domain.User true "account info"
// @Success 200 {string} string "user created"
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.UserService.Create(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created",
	})
}

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.UserService.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
