package handler

import (
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUser()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{
		"users": users,
	})
}

func (h *userHandler) New(c *gin.Context) {

	c.HTML(http.StatusOK, "user_new.html", nil)
}

func (h *userHandler) Create(c *gin.Context) {
	var input user.FormCreateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	registerInput := user.RegisterUserInput{
		Name:       input.Name,
		Email:      input.Email,
		Occupation: input.Occupation,
		Password:   input.Passowrd,
	}

	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/users")
}
