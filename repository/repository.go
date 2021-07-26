package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
	FindById(client *mongo.Client, id int64)
	FindAll(client *mongo.Client, collection string)
	Save(client *mongo.Client, document interface{})
	Delete(client *mongo.Client, id int64)
}
