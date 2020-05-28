package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kadekchrisna/openbook/logger"
	_ "github.com/kadekchrisna/openbook/utils/env"
)

var (
	router = gin.Default()
)

// StartApplication is for Starting Appliclation
func StartApplication() {
	port := ":3080"
	endpoints()
	logger.Info(fmt.Sprintf("Server started at %s", port))
	router.Run(port)
}
