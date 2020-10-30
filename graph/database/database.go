package database

import (
	"context"
	"log"
	"time"

	"github.com/H-Richard/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

// Save   - Save an item into our DB
func (db *DB) Save(input *model.NewDog) *model.Dog {

}

// FindByID  - Find an item by ID
func (db *DB) FindByID(ID string) *model.Dog {

}

// GetAll  - Get all Items from the DB
func (db *DB) GetAll() []*model.Dog {

}
