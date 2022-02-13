package repository

import (
	"github.com/pedroribeiro/starwars-api/internal/domain/entity"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlantesRespositoryMock struct {
	mock.Mock
}

func (m *PlantesRespositoryMock) ListPlanets() (*[]entity.Planet, error) {
	args := m.Called()
	return args.Get(0).(*[]entity.Planet), args.Error(1)
}

func (m *PlantesRespositoryMock) GetPlanetByName(name string) (*entity.Planet, error) {
	args := m.Called(name)
	return args.Get(0).(*entity.Planet), args.Error(1)
}

func (m *PlantesRespositoryMock) GetPlanetById(id string) (*entity.Planet, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Planet), args.Error(1)
}

func (m *PlantesRespositoryMock) InsertPlanet(p *entity.Planet) (*mongo.InsertOneResult, error) {
	args := m.Called(p)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *PlantesRespositoryMock) DeletePlanetById(id string) (*mongo.DeleteResult, error) {
	args := m.Called(id)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}
