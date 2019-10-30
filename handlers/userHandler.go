package handlers

import (
	"crud/bindings"
	"crud/helpers"
	"crud/models"
	"fmt"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var response helpers.ResponseHelper

type UserHandler struct {
	DbSession *mgo.Session
	DbName    string
}

func (h *UserHandler) HelloWorld(c echo.Context) error {
	return response.SendSuccess(c, "success", response.EmptyJsonMap())
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	var users []models.User

	ds := h.DbSession.Copy()
	defer ds.Close()
	err := ds.DB(h.DbName).C("users").Find(nil).All(&users)
	if err != nil {
		return response.SendError(c, err.Error(), response.EmptyJsonMap)
	}
	return response.SendSuccess(c, "success", users)
}

func (h *UserHandler) PostUser(c echo.Context) error {
	var (
		err       error
		input     bindings.User
		modelUser models.User
	)

	err = c.Bind(&input)
	if err != nil {
		return response.SendBadRequest(c, err.Error(), response.EmptyJsonMap)
	}

	ds := h.DbSession.Copy()
	defer ds.Close()
	table := ds.DB(h.DbName).C("users")

	modelUser = models.User{
		Id:       bson.NewObjectId(),
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
	}

	// create new member
	err = table.Insert(modelUser)
	if err != nil {
		return response.SendError(c, err.Error(), input)
	}
	return response.SendSuccess(c, "success", modelUser)
}

func (h *UserHandler) PutUser(c echo.Context) error {
	var (
		err       error
		input     bindings.User
		modelUser models.User
	)

	err = c.Bind(&input)
	if err != nil {
		return response.SendBadRequest(c, err.Error(), response.EmptyJsonMap)
	}

	id := c.Param("id")

	ds := h.DbSession.Copy()
	defer ds.Close()
	table := ds.DB(h.DbName).C("users")

	err = table.Update(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": bson.M{
			"email":    input.Email,
			"password": input.Password,
			"name":     input.Name,
		}},
	)

	if err != nil {
		return response.SendError(c, err.Error(), input)
	}

	err = table.FindId(bson.ObjectIdHex(id)).One(&modelUser)
	if err != nil {
		return response.SendError(c, err.Error(), input)
	}
	return response.SendSuccess(c, "success", modelUser)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	var (
		err       error
		modelUser models.User
	)

	id := c.Param("id")
	fmt.Println(id)

	ds := h.DbSession.Copy()
	defer ds.Close()
	table := ds.DB(h.DbName).C("users")
	err = table.FindId(bson.ObjectIdHex(id)).One(&modelUser)
	if err != nil {
		fmt.Println(err.Error())
		return response.SendError(c, err.Error(), response.EmptyJsonMap())
	}
	return response.SendSuccess(c, "success", modelUser)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var (
		err error
	)

	id := c.Param("id")
	fmt.Println(id)

	ds := h.DbSession.Copy()
	defer ds.Close()
	table := ds.DB(h.DbName).C("users")
	err = table.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		fmt.Println(err.Error())
		return response.SendError(c, err.Error(), response.EmptyJsonMap())
	}
	return response.SendSuccess(c, "success", response.EmptyJsonMap())
}
