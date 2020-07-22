package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
	"iris-gorm/database"
	"iris-gorm/model"
)

func main() {
	db := database.Mydb()
	//defer close db
	defer func() {
		_ = db.Close()
	}()

	// Run the function to create the new Iris App
	app := newApp(db)

	// Start the web server on port 8080
	app.Run(iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		//Iris-config.yaml
		iris.WithConfiguration(
			iris.YAML("./conf/config.yaml"), //loading config.yaml
		))
}

func newApp(db *gorm.DB) *iris.Application {
	// Initialize a new Iris App
	app := iris.New()
	// Register the request handler for the endpoint "/"
	app.Get("/", func(ctx iris.Context) {
		// Return something by adding it to the context
		ctx.Text("Hello World")
	})
	app.Get("/ping", func(ctx iris.Context) {
		// Return something by adding it to the context
		ctx.HTML("pong")
	})
	// Register an endpoint with a variable
	app.Get("/tag/{name:string}", func(ctx iris.Context) {
		ctx.Text(fmt.Sprintf("Hello %s", ctx.Params().Get("name")))
	})
	// Define the slice for the result
	var movices []model.Movie
	// Endpoint to perform the database request
	app.Get("/movie", func(ctx iris.Context) {
		//gorm set sql
		db.Order("Id desc").Limit(1).Find(&movices)
		_, _ = ctx.JSON(movices)
	})
	return app
}
