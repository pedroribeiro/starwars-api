package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/pedroribeiro/starwars-api/internal/domain/entity"
	"github.com/pedroribeiro/starwars-api/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlanetsUseCase struct {
	Repo repository.IPlanetsRepository
}

func (u *PlanetsUseCase) ListPlanets() (*[]entity.Planet, error) {

	p, err := u.Repo.ListPlanets()

	if err != nil {
		return nil, err
	}

	ch := make(chan entity.Planet)
	var wg sync.WaitGroup

	res := []entity.Planet{}

	wg.Add(len(*p))

	for _, item := range *p {
		go func(item entity.Planet) {
			getMoviesCount(&item, ch, &wg)
		}(item)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		res = append(res, r)
	}

	return &res, nil
}

func (u *PlanetsUseCase) GetPlanetByName(name string) (*entity.Planet, error) {
	p, err := u.Repo.GetPlanetByName(name)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	ch := make(chan entity.Planet)

	wg.Add(1)

	go getMoviesCount(p, ch, &wg)

	if err != nil {
		return nil, err
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		p = &r
	}

	return p, nil
}

func (u *PlanetsUseCase) GetPlanetById(id string) (*entity.Planet, error) {
	p, err := u.Repo.GetPlanetById(id)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	ch := make(chan entity.Planet)

	wg.Add(1)

	go getMoviesCount(p, ch, &wg)

	if err != nil {
		return nil, err
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		p = &r
	}

	return p, nil
}

func (u *PlanetsUseCase) InsertPlanet(p *entity.Planet) (*mongo.InsertOneResult, error) {
	return u.Repo.InsertPlanet(p)
}

func (u *PlanetsUseCase) DeletePlanetById(id string) (*mongo.DeleteResult, error) {
	return u.Repo.DeletePlanetById(id)
}

func getMoviesCount(p *entity.Planet, ch chan<- entity.Planet, wg *sync.WaitGroup) {
	defer wg.Done()

	endpoint := fmt.Sprintf("https://swapi.dev/api/planets/?search=%s", strings.Fields(p.Name)[0])

	res, err := http.Get(endpoint)

	if err != nil {
		log.Println(err)
	}

	payload := &entity.PlanetsPayload{}

	json.NewDecoder(res.Body).Decode(payload)

	defer res.Body.Close()

	if len(payload.Results) > 0 {
		p.Movies_count = len(payload.Results[0].Films)
	} else {
		p.Movies_count = 0
	}

	ch <- *p
}
