package main

import (
	"crud/configs/boot"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	handler := boot.HTTPHandler{E: e}

	handler.RegisterApiHandler()

	e.Logger.Fatal(e.Start(":1323"))
}
