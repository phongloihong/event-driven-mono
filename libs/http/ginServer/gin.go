package ginServer

import (
	"github.com/gin-gonic/gin"
)

type GinConfig struct {
	Port string `mapstructure:"port" validate:"required"`
}

func NewHttpServer(cfg *GinConfig) *gin.Engine {
	r := gin.Default()
	return r
}
