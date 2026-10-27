package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"quiz.com/quiz/internal/collection"
	"quiz.com/quiz/internal/controller"
	"quiz.com/quiz/internal/service"
)

type App struct {
	httpServer *fiber.App
	database   *mongo.Database

	quizService  *service.QuizService
	netService   *service.NetService
	imageService *service.ImageService
}

func (a *App) Init() {
	a.setupDb()
	a.setupServices()
	a.setupHttp()

	log.Fatal(a.httpServer.Listen(":3000"))
}

func (a *App) setupHttp() {
	app := fiber.New()
	app.Use(cors.New())

	quizController := controller.Quiz(a.quizService)
	app.Get("/api/quizzes", quizController.GetQuizzes)
	app.Post("/api/quizzes", quizController.CreateQuiz)
	app.Get("/api/quizzes/:quizId", quizController.GetQuizById)
	app.Put("/api/quizzes/:quizId", quizController.UpdateQuizById)

	imageController := controller.Image(a.imageService)
	app.Post("/api/upload", imageController.UploadImage)
	app.Post("/api/images", imageController.UploadImage)
	app.Get("/api/images/:id", imageController.GetImage)
	app.Delete("/api/images/:id", imageController.DeleteImage)

	wsController := controller.Ws(a.netService)
	app.Get("/ws", websocket.New(wsController.Ws))

	a.httpServer = app
}

func (a *App) setupServices() {
	a.quizService = service.Quiz(collection.Quiz(a.database.Collection("quizzes")))
	a.netService = service.Net(a.quizService)

	imageColl, err := collection.Image(a.database)
	if err != nil {
		panic(err)
	}
	a.imageService = service.Image(imageColl)
}

func (a *App) setupDb() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	mongoURI := os.Getenv("MONGODB_URI")
	fmt.Println("MONGODB_URI:", mongoURI)
	if mongoURI == "" {
		panic("MONGODB_URI environment variable is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	a.database = client.Database("quiz")
}
