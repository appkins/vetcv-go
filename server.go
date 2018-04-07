package main

import (
	"github.com/appkins/vetcv-go/api"
	"github.com/appkins/vetcv-go/models"
	"github.com/appkins/vetcv-go/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
