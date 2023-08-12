package mongo_driver

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Init(databaseName string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Make MongoDB URI
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// fmt.println(dbUser, dbPass, dbHost)
	mongodbURI := "mongodb://mongo:U73R5S6hkwjlEwAy4e3H@containers-us-west-45.railway.app:5758"

	// Set client options
	clientOptions := options.Client().ApplyURI(mongodbURI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	//err = client.Ping(ctx, nil)
	//if err != nil {
	//	panic(err)
	//}

	database := client.Database(databaseName)
	return database
}

func Close(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return db.Client().Disconnect(ctx)
}
