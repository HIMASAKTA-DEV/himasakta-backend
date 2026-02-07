package response

import (
	"net/http"

	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"-"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Error      any    `json:"error,omitempty"`
	Data       any    `json:"data,omitempty"`
	Meta       any    `json:"meta,omitempty"`
}

func NewSuccess(msg string, data any, meta ...any) Response {
	res := Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    msg,
		Data:       data,
	}

	if len(meta) > 0 {
		res.Meta = meta[0]
	}

	return res
}

// response default status code internal server error (500)
func NewFailed(msg string, err error, data ...any) Response {
	res := Response{
		StatusCode: http.StatusInternalServerError,
		Success:    false,
		Message:    msg,
		Error:      err.Error(),
	}

	if myErr, ok := err.(myerror.Error); ok {
		res.StatusCode = myErr.StatusCode
	}

	if len(data) > 0 {
		res.Data = data
	}

	return res
}

func NewFailedWithCode(statusCode int, msg string, err error, data ...any) Response {
	res := Response{
		StatusCode: statusCode,
		Success:    false,
		Message:    msg,
		Error:      err.Error(),
	}

	if len(data) > 0 {
		res.Data = data
	}

	return res
}

func (r Response) ChangeStatusCode(statusCode int) Response {
	res := r
	res.StatusCode = statusCode
	return res
}

func (r Response) Send(ctx *gin.Context) {
	sendStatus := r.StatusCode
	ctx.JSON(sendStatus, r)
}

func (r Response) SendWithAbort(ctx *gin.Context) {
	sendStatus := r.StatusCode
	ctx.AbortWithStatusJSON(sendStatus, r)
}
