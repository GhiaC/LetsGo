package main

import (
	"context"
	"github.com/aoacademy/letsgo/conf"
	"github.com/aoacademy/letsgo/database"
	"github.com/aoacademy/letsgo/router"
)

func main() {
	configuration := conf.New()
	ctx := context.Background()
	client := database.New(ctx, configuration)
	collection := database.GetCollection(client, configuration.Database, configuration.Collection)
	routing := router.AddPaths(collection, ctx)
	router.BindHttp(routing, configuration)
}
