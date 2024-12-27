package _type

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Router          *gin.Engine
	Database        *mongo.Database
	MangaCollection *mongo.Collection
}

type Manga struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string
	MaxPage int `bson:"max_page"`
}
