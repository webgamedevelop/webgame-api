package validator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"k8s.io/apimachinery/pkg/api/resource"
)

func RegisterValidation() error {
	var (
		v  *validator.Validate
		ok bool
	)
	if v, ok = binding.Validator.Engine().(*validator.Validate); !ok {
		return fmt.Errorf("type assertion failed")
	}
	if err := v.RegisterValidation("k8sCpu", k8sCpuValidator); err != nil {
		return err
	}
	if err := v.RegisterValidation("k8sMemory", k8sMemoryValidator); err != nil {
		return err
	}
	return nil
}

var k8sCpuValidator validator.Func = func(fl validator.FieldLevel) bool {
	if r, err := resource.ParseQuantity(fl.Field().String()); err != nil {
		return false
	} else {
		tmp := r.String()
		if strings.HasSuffix(tmp, "m") {
			tmp = strings.TrimSuffix(tmp, "m")
		}
		if _, err := strconv.Atoi(tmp); err != nil {
			return false
		}
		return true
	}
}

var k8sMemoryValidator validator.Func = func(fl validator.FieldLevel) bool {
	if r, err := resource.ParseQuantity(fl.Field().String()); err != nil {
		return false
	} else {
		tmp := r.String()
		if strings.HasSuffix(tmp, "M") || strings.HasSuffix(tmp, "Mi") || strings.HasSuffix(tmp, "G") || strings.HasSuffix(tmp, "Gi") {
			return true
		}
		return false
	}
}
