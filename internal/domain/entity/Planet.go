package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Planet struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" binding:"required"`
	Climate      string             `json:"climate" binding:"required"`
	Terrain      string             `json:"terrain" binding:"required"`
	Movies_count int                `json:"movies_count"`
}

type ApiPlanet struct {
	Films   []string `json:"films"`
	Name    string   `json:"name" binding:"required"`
	Climate string   `json:"climate" binding:"required"`
	Terrain string   `json:"terrain" binding:"required"`
}

type PlanetsPayload struct {
	Results []ApiPlanet `json:"results"`
}
