package pokemon

import (
	"workshop-go-gin-basic/api/pokemon/controllers"
	"workshop-go-gin-basic/api/pokemon/models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func createTable(db *pg.DB) {
	err := db.Model((*models.Pokemon)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

// Init init pokemon api
func Init(engine *gin.Engine, db *pg.DB) {
	createTable(db)
	router := engine.Group("/pokemon")
	{
		router.GET(":pokemonId", controllers.Get)
		router.POST("", controllers.Create)
		router.GET("", controllers.List)
		router.PUT(":pokemonId", controllers.Update)
		router.DELETE(":pokemonId", controllers.Delete)
	}
}
