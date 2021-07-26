package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AutorEntity struct {
	Id        primitive.ObjectID `bson:"_id"`
	Nome      string             `bson:"nome"`
	Email     string             `bson:"email"`
	Descricao string             `bson:"descricao"`
	Instante  time.Time          `bson:"instante"`
}

func getCollection(client mongo.Client, database string, collection string) *mongo.Collection {
	clientCollection := client.Database(database).Collection(collection)
	return clientCollection
}

func (autor AutorEntity) Save(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := getCollection(*client, "casadocodigo", "autor")
	_, err := col.InsertOne(ctx, autor)

	return err
}
