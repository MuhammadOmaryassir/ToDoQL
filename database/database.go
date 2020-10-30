package database

import (
	"context"
	"log"
	"time"

	"github.com/Omaroovee/ToDoQL/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

// Connect  - Connect to the DB
func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://<dbuser>:<dbpassword>@ds029979.mlab.com:29979/todoql"))
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
func (db *DB) Save(input *model.NewTodo) *model.Todo {
	collection := db.client.Database("todoql").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Todo{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		Description: input.Description,
		Status:      input.Status,
	}
}

// FindByID  - Find an item by ID
func (db *DB) FindByID(ID string) *model.Todo {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("todoql").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	todo := model.Todo{}
	res.Decode(&Todo)
	return &todo
}

// GetAll  - Get all Items from the DB
func (db *DB) GetAll() []*model.Todo {
	collection := db.client.Database("todoql").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var todos []*model.Todo
	for cur.Next(ctx) {
		var todo *model.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	return todos
}
