package main

import (
	"datadog-userpost/src/datadog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	YourAPIKey := "fe5f5e7807f1c404578370cb78ee9f92"
	YourApplicationKey := "4dcb0a60eb4554c8d89e8320b9da9937d515e8fa"

	datadogDomain := datadog.NewDomainImplementation(YourAPIKey, YourApplicationKey)
	datadogUsecase := datadog.NewUsecaseImplementation(datadogDomain)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.POST("/create_user", func(c *gin.Context) {
		var reqData datadog.AddUserRequest
		if c.BindJSON(&reqData) == nil {
			respData, err := datadogUsecase.CreateUser(reqData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"data": respData,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": respData,
			})
			return
		}
	})
	r.POST("/disable_user", func(c *gin.Context) {
		var reqData datadog.UserHandleRequest
		if c.BindJSON(&reqData) == nil {
			respData, err := datadogUsecase.DisableUser(reqData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"data": respData,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": respData,
			})
			return
		}
	})
	r.POST("/get_users", func(c *gin.Context) {

		respData, err := datadogUsecase.GetAllUser()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": respData,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": respData,
		})
		return
	})
	r.POST("/get_user", func(c *gin.Context) {
		var reqData datadog.UserHandleRequest
		if c.BindJSON(&reqData) == nil {
			respData, err := datadogUsecase.GetUser(reqData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"data": respData,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": respData,
			})
			return
		}
	})
	r.POST("/update_user", func(c *gin.Context) {
		var reqData datadog.UpdateUserDatadogRequest
		if c.BindJSON(&reqData) == nil {
			respData, err := datadogUsecase.UpdateUser(reqData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"data": respData,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": respData,
			})
			return
		}
	})

	r.Run()

}
