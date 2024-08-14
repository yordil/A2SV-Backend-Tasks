package bootstrap

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	fmt.Println("here worked")
	return *app
}

// func (app *Application) CloseDBConnection() {
// 	CloseMongoDBConnection(app.Mongo)
// }
