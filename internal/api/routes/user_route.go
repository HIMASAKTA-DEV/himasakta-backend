package routes

// func User(app *gin.Engine, usercontroller controller.UserController, middleware middleware.Middleware) {
// 	routes := app.Group("/api/v1/user")
// 	{
// 		routes.POST("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleSuperAdmin)), usercontroller.Create)
// 		routes.GET("", middleware.Authenticate(), usercontroller.GetAll)
// 		routes.GET("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleSuperAdmin)), usercontroller.GetById)
// 		routes.PUT("/:id", middleware.Authenticate(), usercontroller.Update)
// 		routes.DELETE("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleSuperAdmin)), usercontroller.Delete)
// 	}
// }
