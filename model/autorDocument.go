package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Autor struct {
	Id        primitive.ObjectID `bson:"_id"`
	Nome      string             `bson:"nome"`
	Email     string             `bson:"email"`
	Descricao string             `bson:"descricao"`
	Instante  time.Time          `bson:"instante"`
}
