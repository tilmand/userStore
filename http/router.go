package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func configureRouter(api *api) *gin.Engine {
	router := gin.Default()

	router.Use(api.auth.Validate)

	router.GET("/profile", api.Users().GetAllProfiles)
	router.GET("/profile/:username", api.Users().GetProfileByUsername)

	router.NoRoute(func(c *gin.Context) {
		log.Println("route not found")
		c.JSON(http.StatusNotFound, errors.New("route not found"))
	})

	return router
}
