package helper

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Response level
const (
	WARNING = "warning"
	SUCCESS = "success"
	INFO    = "info"
	DANGER  = "danger"
	ERROR   = "error"
)

type (
	// Payload ...
	Payload struct {
		Message string `json:"message"`
	}

	// Response ...
	Response struct {
		State   string      `json:"state"`
		Code    int         `json:"code"`
		Payload interface{} `json:"payload"`
		Message interface{} `json:"message"`
	}
)

// SetMessage ...
func (r *Response) SetMessage(message interface{}) *Response {
	r.Message = message
	return r
}

// SetPayload ...
func (r *Response) SetPayload(payload interface{}) *Response {
	r.Payload = payload
	return r
}

// SetStatusCode ...
func (r *Response) SetStatusCode(code int) *Response {
	r.Code = code
	return r
}

// SetState ...
func (r *Response) SetState(state string) *Response {
	r.State = state
	return r
}

// JSON ...
func (r *Response) JSON(c echo.Context) error {
	return c.JSON(http.StatusOK, r)
}

func ResponseSuccess(c echo.Context, payload interface{}) error {
	resp := &Response{}
	return resp.SetState(SUCCESS).
		SetStatusCode(http.StatusOK).
		SetPayload(payload).
		JSON(c)
}

func MessageSuccess(c echo.Context, message string) error {
	resp := &Response{}
	return resp.SetState(SUCCESS).
		SetStatusCode(http.StatusOK).
		SetMessage(message).
		JSON(c)
}
