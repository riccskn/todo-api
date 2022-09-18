package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"todo-api/controller"
)

func Initialize(engine *gin.Engine, db *gorm.DB) {

	api := engine.Group("/api")
	{
		group := api.Group("/v1")
		{

			ctrl := controller.New(db)

			group.GET("/todos", ctrl.All)
			todo := group.Group("/todos")
			{
				todo.POST("/", ctrl.Create)
				todo.GET("/:id", ctrl.Get)
				todo.DELETE("/:id", ctrl.Del)
				todo.PUT("/:id", ctrl.Update)
			}
		}
	}

	log.Println("INFO: successfully connected to database")
}
