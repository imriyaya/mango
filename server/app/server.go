package app

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imriyaya/mango/app/handlers"
	"github.com/imriyaya/mango/app/handlers/manga"
	"github.com/imriyaya/mango/app/handlers/manga/chapter/page"
	"github.com/imriyaya/mango/app/type"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Start() {
	loadEnv()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	pingErr := client.Ping(context.TODO(), nil)
	if pingErr != nil {
		panic(pingErr)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	server := _type.Server{
		Router:          router,
		Database:        client.Database("mango"),
		MangaCollection: client.Database("mango").Collection("mangas"),
	}

	handlers.StatusHandler(&server)
	manga.Get(&server)
	manga.List(&server)
	page.Get(&server)
	router.Run()
}
