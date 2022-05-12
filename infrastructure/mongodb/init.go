package mongodb

import (
	"context"
	"fmt"
	"github.com/SananGuliyev/gossignment/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewStorage(config *config.MongoDbConfig) *mongo.Database {
	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		config.Username,
		config.Password,
		config.Host,
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}

	return client.Database(config.Database)
}
