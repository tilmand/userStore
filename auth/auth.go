package auth

import (
	"log"
	"net/http"
	"userStore/model"
	"userStore/store"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	store *store.Store
}

type AuthMiddleware interface {
	Validate(c *gin.Context)
}

func NewAuthMiddleware(store *store.Store) *Middleware {
	var middleware = &Middleware{
		store: store,
	}

	return middleware
}

func (m Middleware) Validate(c *gin.Context) {
	ctx := c.Request.Context()
	apiKey := c.GetHeader("Api-key")
	if apiKey == "" {
		log.Println("no api-key header")
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrInvalidBody)

		return
	}

	count, err := m.store.AuthRepository.Find(apiKey, ctx)
	if err != nil {
		log.Println("apiKey not found:", err)
		c.AbortWithStatusJSON(http.StatusForbidden, model.ErrForbidden)

		return
	}

	if count == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, model.ErrForbidden)

		return
	}
}
