package helpers

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResponseHelper struct {
}

//response format
type responseFormat struct {
	C        echo.Context
	Status   string
	Message  string
	Data     interface{}
	Code     int
	CodeType string
}

//set response
func (h *ResponseHelper) SetResponse(c echo.Context, status string, message string, data interface{}, code int, codeType string) responseFormat {
	return responseFormat{c, status, message, data, code, codeType}
}

func (h *ResponseHelper) SendResponse(res responseFormat) error {

	if len(res.Message) == 0 {
		res.Message = `success`
	}

	var resCode int
	if res.Code != 200 {
		resCode = http.StatusBadRequest
	} else {
		resCode = http.StatusOK
	}

	return res.C.JSON(resCode, map[string]interface{}{
		"code":      res.Code,
		"code_type": res.CodeType,
		"message":   res.Message,
		"data":      res.Data,
	})
}

func (u *ResponseHelper) EmptyJsonMap() map[string]interface{} {
	return make(map[string]interface{})
}

// TODO: Send success response to consumers.
func (u *ResponseHelper) SendSuccess(c echo.Context, message string, data interface{}) error {
	res := u.SetResponse(c, "success", message, data, 200, `success`)

	return u.SendResponse(res)
}

// TODO: Send bad request response to consumers.
func (u *ResponseHelper) SendBadRequest(c echo.Context, message string, data interface{}) error {
	res := u.SetResponse(c, "error", message, data, 400, `badRequest`)
	return u.SendResponse(res)
}

func (u *ResponseHelper) SendError(c echo.Context, message string, data interface{}) error {
	res := u.SetResponse(c, "error", message, data, 500, `internalServerError`)
	return u.SendResponse(res)
}
