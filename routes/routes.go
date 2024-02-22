package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(ItemRoute *gin.Engine) {
	ItemRoute.GET("/items", controllers.GetItems())
	ItemRoute.GET("/items/:id", controllers.GetItem())
	ItemRoute.POST("/items/create", controllers.CreateItem())
	ItemRoute.PUT("/items/update/:id", controllers.UpdateItem())
	ItemRoute.DELETE("/items/delete/:id", controllers.DeleteItem())
}

func UserRoutes(UserRoute *gin.Engine) {
	UserRoute.GET("/users", controllers.GetUser())
	UserRoute.POST("/users/create", controllers.CreateUser())
	UserRoute.POST("/users/signup", controllers.UserSignup())
	UserRoute.POST("/users/login", controllers.UserLogin())
	UserRoute.PUT("/users/update/:id", controllers.UpdateUser())
	UserRoute.DELETE("/users/delete/:id", controllers.DeleteUser())
}
