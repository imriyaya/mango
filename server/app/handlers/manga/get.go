package manga

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	_type "github.com/imriyaya/mango/app/type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Get(server *_type.Server) {
	server.Router.GET("/manga/:manga_id", func(c *gin.Context) {
		mangaId, objectIdError := primitive.ObjectIDFromHex(c.Param("manga_id"))
		if objectIdError != nil {
			_ = c.AbortWithError(http.StatusBadRequest, objectIdError)
			return
		}

		var manga bson.M
		err := server.MangaCollection.FindOne(context.TODO(), bson.D{{"_id", mangaId}}).Decode(&manga)

		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, manga)
	})
}
