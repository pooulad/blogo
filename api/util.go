package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamByName(ctx *gin.Context, paramName string) (interface{}, error) {
	param := ctx.Param(paramName)

	if param == "" {
		return nil, fmt.Errorf("param is empty")
	}

	if id, err := strconv.Atoi(param); err == nil {
		return id, nil
	}

	return param, nil
}
