package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func badResponse(c *gin.Context, code int, err ...error) {
	errs := errors.Join(err...)
	klog.Error(errs)
	response := Response(&simpleResponse{}, Code(code), Message(errs.Error()))
	c.JSON(code, response)
}

func okResponse(c *gin.Context, data any) {
	response := Response(&simpleResponse{}, Code(http.StatusOK), Message("success"), &ResponseExtend{Data: data})
	c.JSON(http.StatusOK, response)
}
