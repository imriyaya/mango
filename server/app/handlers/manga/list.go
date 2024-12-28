package manga

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/imriyaya/mango/app/type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

func List(server *_type.Server) {
	server.Router.GET("/manga/list", func(c *gin.Context) {
		skip, err := strconv.Atoi(c.Request.URL.Query().Get("skip"))
		if err != nil {
			c.JSON(400, err)
			return
		}

		count, err := server.MangaCollection.EstimatedDocumentCount(context.TODO())
		if err != nil {
			c.JSON(500, err)
			return
		}

		opts := options.Find().SetLimit(30).SetSkip(int64(skip)).SetProjection(bson.D{{"chapters", 0}})
		cursor, err := server.MangaCollection.Find(context.TODO(), bson.D{}, opts)
		if err != nil {
			c.JSON(500, err)
			return
		}

		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			c.JSON(500, err)
			return
		}

		if results == nil {
			results = make([]bson.M, 0)
		}

		c.JSON(200, gin.H{
			"count":  count,
			"result": results,
		})
	})
}
