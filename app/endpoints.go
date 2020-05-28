package app

import (
	"github.com/kadekchrisna/openbook/controllers/ping"
	"github.com/kadekchrisna/openbook/controllers/users"
)

func endpoints() {
	router.GET("/ping", ping.Ping)

	router.POST("/user", users.CreateUser)
	router.POST("/user/auth", users.Login)

	router.GET("/user/:user_id", users.GetUser)
	router.GET("/internal/user/search", users.Search)

	router.PUT("/user/:user_id", users.UpdateUser)
	router.PATCH("/user/:user_id", users.UpdateUser)

	router.DELETE("/user/:user_id", users.DeleteUser)

}
