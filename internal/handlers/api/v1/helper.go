package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func BadResponse(c *gin.Context, code int, err ...error) {
	errs := errors.Join(err...)
	klog.Error(errs)
	response := SimpleResponse(&simpleResponse{}, Code(code), Message(errs.Error()))
	c.JSON(code, response)
}

func EmptyResponse(c *gin.Context, code int) {
	response := SimpleResponse(&simpleResponse{}, Code(code), Message("success"))
	c.JSON(code, response)
}

func DetailResponse[T any](c *gin.Context, data T) {
	response := &detailResponse[T]{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func ListResponse[T []E, E any](c *gin.Context, data T) {
	response := &listResponse[T, E]{
		Code:    http.StatusOK,
		Message: "success",
		Len:     len(data),
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}
