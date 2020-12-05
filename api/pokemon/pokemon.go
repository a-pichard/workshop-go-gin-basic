package pokemon

import (
	"workshop-go-gin-basic/api/pokemon/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func Init(engine *gin.Engine, db *pg.DB) {
	router := engine.Group("/pokemon")
	{
		router.GET("", controllers.Get)
	}
}
