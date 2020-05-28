package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kadekchrisna/openbook-oauth-go/oauth"
	"github.com/kadekchrisna/openbook/domains/users"
	"github.com/kadekchrisna/openbook/services"
	"github.com/kadekchrisna/openbook/utils/errors"
)

// GetUserId parsing id on param URL returning int64 or error
func GetUserId(idParam string) (int64, *errors.ResErr) {
	userId, errParse := strconv.ParseInt(idParam, 10, 64)
	if errParse != nil {
		error := errors.NewBadRequestError("invalid id format")
		return 0, error
	}
	return userId, nil
}

// GetUser is for getting user by id
func GetUser(c *gin.Context) {

	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	userId, errParse := GetUserId(c.Param("user_id"))
	if errParse != nil {
		c.JSON(errParse.Status, errParse)
		return
	}

	result, errCreate := services.UsersService.GetUser(userId)
	if errCreate != nil {
		// TODO Handle error
		c.JSON(errCreate.Status, errCreate)
		return
	}

	if oauth.GetCallerId(c.Request) == userId {
		c.JSON(http.StatusOK, result.Marshall(false))
		return
	}
	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
	return

}

// CreateUser is for creating new user
func CreateUser(c *gin.Context) {
	var user users.User
	// From line 23-38 is the same functionality as line 40
	//
	// bytes, readErr := ioutil.ReadAll(c.Request.Body)
	// if readErr != nil {
	// 	// TODO Handle error
	// 	// log.Fatalf("Error occurred when reading body. \n %s", readErr)
	// 	fmt.Printf("Error occurred when reading body. \n %s\n", readErr.Error())
	// 	// c.String(http.StatusBadRequest, "Help")
	// 	return
	// }
	// unmarshErr := json.Unmarshal(bytes, &user)
	// if unmarshErr != nil {
	// 	// TODO Handle error
	// 	// log.Fatalf("Error occurred when unmarshalling body. \n %s", unmarshErr)
	// 	// fmt.Printf("Error occurred when unmarshalling body. \n %s\n", unmarshErr.Error())
	// 	c.String(http.StatusBadRequest, "Help")
	// 	return
	// }

	err := c.ShouldBindJSON(&user)
	if err != nil {
		// TODO Handle error
		fmt.Printf("Error occurred when unmarshalling body. \n %s\n", err.Error())
		resErr := errors.NewBadRequestError("invalid json format")
		c.JSON(resErr.Status, resErr)
		return
	}

	result, errCreate := services.UsersService.CreateUser(user)
	if errCreate != nil {
		// TODO Handle error
		// fmt.Printf("Error occurred when unmarshalling body. \n %s\n", errCreate.Error())
		// c.String(http.StatusBadRequest, "Help")
		// return

		c.JSON(errCreate.Status, errCreate)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
	return

}

// UpdateUser updating pusposes
func UpdateUser(c *gin.Context) {
	var user users.User
	userId, errParse := GetUserId(c.Param("user_id"))
	if errParse != nil {
		c.JSON(errParse.Status, errParse)
		return
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		// TODO Handle error
		fmt.Printf("Error occurred when unmarshalling body. \n %s\n", err.Error())
		resErr := errors.NewBadRequestError("invalid json format")
		c.JSON(resErr.Status, resErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	user.Id = userId
	result, errUpdateUser := services.UsersService.UpdateUser(isPartial, user)
	if errUpdateUser != nil {
		c.JSON(errUpdateUser.Status, errUpdateUser)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
	return
}

// DeleteUser deleting user
func DeleteUser(c *gin.Context) {
	userID, errParse := GetUserId(c.Param("user_id"))
	if errParse != nil {
		c.JSON(errParse.Status, errParse)
		return
	}

	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "OK"})
	return
}

// Search user
func Search(c *gin.Context) {
	search := c.Query("search")
	result, errSearch := services.UsersService.SearchUser(search)
	if errSearch != nil {
		c.JSON(errSearch.Status, errSearch)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
	return
}

// Login user
func Login(c *gin.Context) {
	var user users.LoginRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		errJSON := errors.NewBadRequestError("Invalid json format")
		c.JSON(errJSON.Status, errJSON)
		return
	}

	result, err := services.UsersService.LogginUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
	return

}
