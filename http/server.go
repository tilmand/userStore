package http

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"userStore/auth"
	"userStore/store"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

type api struct {
	auth   auth.AuthMiddleware
	sql    *store.Store
	router *gin.Engine

	usersHandler *UsersHandler
}

func NewServer(auth auth.AuthMiddleware, sql *store.Store, port *int) *api {
	api := &api{
		auth: auth,
		sql:  sql,
	}

	api.router = configureRouter(api)
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	server := http.Server{
		Addr:    addr,
		Handler: api.router,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.ListenAndServe()
		if err != nil {
			log.Println("NewServer ListenAndServe err: ", err)
		}
	}()

	return nil
}

func Wait() {
	wg.Wait()
}

func (a *api) Users() *UsersHandler {
	if a.usersHandler == nil {
		a.usersHandler = NewUsersHandler(a)
	}

	return a.usersHandler
}
