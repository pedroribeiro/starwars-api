package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pedroribeiro/starwars-api/internal/config"
	"github.com/pedroribeiro/starwars-api/internal/delivery/api"
	"github.com/pedroribeiro/starwars-api/internal/repository"
	"github.com/pedroribeiro/starwars-api/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client := initDB()

	router := gin.Default()

	// Planet Domain
	planetsCollection := client.Database("starwars").Collection("planets")
	planetRepo := repository.PlanetsRepository{DB: planetsCollection}
	planetUsecase := usecase.PlanetsUseCase{Repo: &planetRepo}
	planetHandler := api.PlanetHandler{PlanetsUseCase: &planetUsecase}
	planetHandler.InitPlanetRoutes(router)

	router.Run(config.LoadEnv("PORT"))

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func initDB() *mongo.Client {

	DB_URL := config.LoadEnv("DB_URL")

	clientOptions := options.Client().ApplyURI(DB_URL)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

	log.Printf("Connected to database!")

	return client
}
