package response

import (
	"errors"
	"net/http"
	"strings"

	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func NewSuccessCreated(msg string, data any, meta ...any) Response {
	res := Response{
		StatusCode: http.StatusCreated,
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
	statusCode := http.StatusInternalServerError
	errStr := myerror.ParseValidationError(err)

	if myErr, ok := err.(myerror.Error); ok {
		statusCode = myErr.StatusCode
	} else if errors.Is(err, gorm.ErrRecordNotFound) || strings.Contains(strings.ToLower(err.Error()), "not found") {
		statusCode = http.StatusNotFound
	}

	res := Response{
		StatusCode: statusCode,
		Success:    false,
		Message:    msg,
		Error:      errStr,
	}

	if len(data) > 0 {
		res.Data = data
	}

	return res
}

func NewFailedWithCode(statusCode int, msg string, err error, data ...any) Response {
	errStr := myerror.ParseValidationError(err)
	res := Response{
		StatusCode: statusCode,
		Success:    false,
		Message:    msg,
		Error:      errStr,
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
