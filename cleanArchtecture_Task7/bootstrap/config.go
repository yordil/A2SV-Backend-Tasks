package bootstrap

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB(){
//    gin.SetMode(gin.ReleaseMode)
  
//   if err != nil {
//     log.Fatal("Error loading .env file")
//   }

  mongoURI := "mongodb+srv://yordi:123456taskmanager@cluster0.4iymqyp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
  if mongoURI == "" {
    log.Fatal("MONGO_URI environment variable not set")
  }

  clientOptions := options.Client().ApplyURI(mongoURI)

  client, err := mongo.Connect(context.TODO(), clientOptions)

  if err != nil {
    log.Fatal(err)
  }

  err = client.Ping(context.TODO(), nil)

  if err != nil {
    log.Fatal(err)
  }
  
  log.Println("Connected to MongoDB!")
  Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
	  return Client.Database("taskmanager").Collection(collectionName)
}