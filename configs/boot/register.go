package boot

import (
	"crud/configs/db/mongo"
	"crud/handlers"

	"github.com/labstack/echo"
)

type HTTPHandler struct {
	E *echo.Echo
}

func (h *HTTPHandler) RegisterApiHandler() *HTTPHandler {
	db := mongo.Info{
		Hostname: "127.0.0.1",
		Database: "go_crud",
		Username: "",
		Password: "",
	}

	dbConnect, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}

	userHandler := &handlers.UserHandler{
		DbSession: dbConnect,
		DbName:    db.Database,
	}
	h.E.GET("/", userHandler.HelloWorld)
	h.E.GET("/user/:id", userHandler.GetUser)
	h.E.GET("/users", userHandler.GetUsers)
	h.E.POST("/user", userHandler.PostUser)
	h.E.PUT("/user/:id", userHandler.PutUser)
	h.E.DELETE("/user/:id", userHandler.DeleteUser)
	return h
}
