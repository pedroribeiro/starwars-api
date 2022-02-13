package usecase

import (
	"github.com/pedroribeiro/starwars-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlanetsUseCase interface {
	ListPlanets() (*[]entity.Planet, error)
	GetPlanetByName(name string) (*entity.Planet, error)
	GetPlanetById(id string) (*entity.Planet, error)
	InsertPlanet(p *entity.Planet) (*mongo.InsertOneResult, error)
	DeletePlanetById(id string) (*mongo.DeleteResult, error)
}
