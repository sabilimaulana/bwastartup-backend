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
		fmt.Println(err.Error())
		//Under Construction
		return
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{
		"users": users,
	})
}
