package handler

import (
	"net/http"
	"nura-fund/helper"
	"nura-fund/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	// Tangkap input dari user
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// Token, err : = h.jwtService.GenerateToken()
	formatter := user.FormatUser(newUser, "tokentokentokentoken")

	// Handler user
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}