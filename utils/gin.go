package utils

import (
	"net/http"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func Build() *gin.Engine {
	logger, _ := zap.NewProduction()
	app := gin.New()
	app.Use(ZapLogger(logger))
	app.Use(ZapRecovery(logger))
	return app
}

func ParseRequest(c *gin.Context, dst interface{}) error {
	binders := []func(any) error{
		func(obj any) error { return c.ShouldBindUri(obj) },
		func(obj any) error { return c.ShouldBindHeader(obj) },
		func(obj any) error { return c.ShouldBindQuery(obj) },
	}

	for _, binder := range binders {
		if err := binder(dst); err != nil {
			return err
		}
	}

	// check if body should be parsed
	if c.Request.Method != http.MethodGet {
		switch c.ContentType() {
		case binding.MIMEJSON:
			if err := c.ShouldBindJSON(dst); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return err
			}
		}
	}

	// run validation at the very end
	return acquireValidator().Struct(dst)
}

var (
	validatorOnce   sync.Once
	sharedValidator *validator.Validate
)

// acquireValidator initializes the validator once and reuses it.
func acquireValidator() *validator.Validate {
	validatorOnce.Do(func() {
		v := validator.New()
		v.RegisterTagNameFunc(preferredFieldName)
		sharedValidator = v
	})
	return sharedValidator
}

// preferredFieldName chooses the "best" tag for reporting validation errors.
func preferredFieldName(field reflect.StructField) string {
	for _, key := range []string{"json", "form", "uri", "header"} {
		if alias := field.Tag.Get(key); alias != "" {
			return alias
		}
	}
	// fallback to struct field name
	return field.Name
}
