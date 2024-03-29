package http

import (
	"context"
	"log"
	"net/http"
	"time"
	"userStore/model"

	"github.com/gin-gonic/gin"
)

type UsersHandlerInterface interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
}

type UsersHandler struct {
	api *api
}

func NewUsersHandler(a *api) *UsersHandler {
	return &UsersHandler{
		api: a,
	}
}
func (h *UsersHandler) GetAllProfiles(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	results, err := h.api.sql.UsersRepository.GetAll(ctx)
	if err != nil {
		log.Println("GetAll GetAll err: ", err)
		c.JSON(http.StatusInternalServerError, model.ErrInternalServerError)

		return
	}
	answer := map[string]interface{}{
		"items": results,
	}

	c.JSON(http.StatusOK, answer)
}

func (h *UsersHandler) GetProfileByUsername(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	username := c.Param("username")

	item, err := h.api.sql.UsersRepository.Find(username, ctx)
	if err != nil {
		log.Println("Get Find err: ", err)
		c.JSON(http.StatusInternalServerError, model.ErrInternalServerError)

		return
	}

	answer := map[string]interface{}{
		"item": item,
	}

	c.JSON(http.StatusOK, answer)
}
