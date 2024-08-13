package bootstrap

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) mongo.Client {
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

mongoURI := env.MongoURI
	
  if mongoURI == "" {
    log.Fatal("MONGO_URI environment variable not set")
  }

  clientOptions := options.Client().ApplyURI(mongoURI)

  client, err := mongo.Connect(ctx, clientOptions)

  if err != nil {
    log.Fatal(err)
  }

  err = client.Ping(context.TODO(), nil)

  if err != nil {
    log.Fatal(err)
  }
  
  log.Println("Connected to MongoDB!")
  
  
  return *client
}

// func CloseMongoDBConnection(client mongo.Client) {
// 	if client == nil {
// 		return
// 	}

// 	err := client.Disconnect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Connection to MongoDB closed.")
// }
