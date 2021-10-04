package handler

import (
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler {
	return &sessionHandler{userService}
}

func (h *sessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}
