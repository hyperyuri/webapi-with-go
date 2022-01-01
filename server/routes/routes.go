package routes

import (
	"github.com/LOO2/business-remote-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		revenues := main.Group("revenue")
		{
			revenues.GET("/", controllers.ShowAllRevenues)
			revenues.GET("/:id", controllers.ShowRevenue)
			revenues.POST("/", controllers.CreateRevenue)
			revenues.PUT("/", controllers.UpdateRevenue)
			revenues.DELETE("/:id", controllers.DeleteRevenue)
		}

		cost_category := main.Group("cost_category")
		{
			cost_category.GET("/", controllers.ShowAllCostCategory)
			cost_category.GET("/:id", controllers.ShowCostCategory)
			cost_category.POST("/", controllers.CreateCostCategory)
			cost_category.PUT("/", controllers.UpdateCostCategory)
			cost_category.DELETE("/:id", controllers.DeleteCostCategory)
		}
	}

	return router
}
