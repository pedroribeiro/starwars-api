package repository

import (
	"context"
	"log"

	"github.com/pedroribeiro/starwars-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlanetsRepository struct {
	DB *mongo.Collection
}

func (r *PlanetsRepository) ListPlanets() (*[]entity.Planet, error) {
	var p []entity.Planet

	cursor, err := r.DB.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = cursor.All(context.TODO(), &p)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &p, nil
}

func (r *PlanetsRepository) GetPlanetByName(name string) (*entity.Planet, error) {
	var p entity.Planet

	err := r.DB.FindOne(context.TODO(), bson.M{"name": name}).Decode(&p)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &p, nil
}

func (r *PlanetsRepository) GetPlanetById(id string) (*entity.Planet, error) {
	var p entity.Planet

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
		return nil, err
	}

	err = r.DB.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&p)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &p, nil
}

func (r *PlanetsRepository) InsertPlanet(p *entity.Planet) (*mongo.InsertOneResult, error) {

	p.Id = primitive.NewObjectID()

	res, err := r.DB.InsertOne(context.TODO(), p)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return res, nil
}

func (r *PlanetsRepository) DeletePlanetById(id string) (*mongo.DeleteResult, error) {

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
		return nil, err
	}

	res, err := r.DB.DeleteOne(context.TODO(), bson.M{"_id": objectId})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return res, nil
}
