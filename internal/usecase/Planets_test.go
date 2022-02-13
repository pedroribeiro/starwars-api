package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/pedroribeiro/starwars-api/internal/domain/entity"
	"github.com/pedroribeiro/starwars-api/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestPlanetsUseCase_ListPlanets(t *testing.T) {
	type fields struct {
		Repo *repository.PlantesRespositoryMock
	}
	repo := new(repository.PlantesRespositoryMock)
	tests := []struct {
		name    string
		fields  fields
		want    *[]entity.Planet
		mock    *[]entity.Planet
		wantErr bool
		err     error
	}{
		{
			name: "ListPlanets",
			fields: fields{
				Repo: repo,
			},
			mock: &[]entity.Planet{
				{
					Name: "Tatooine",
				},
				{
					Name: "Alderaan",
				},
			},
			want: &[]entity.Planet{
				{
					Name:         "Tatooine",
					Movies_count: 5,
				},
				{
					Name:         "Alderaan",
					Movies_count: 2,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "ListPlanets-Error",
			fields: fields{
				Repo: repo,
			},
			mock:    nil,
			want:    nil,
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Repo.On("ListPlanets").Return(tt.mock, tt.err).Once()
			u := &PlanetsUseCase{
				Repo: tt.fields.Repo,
			}

			got, err := u.ListPlanets()

			if (err != nil) != tt.wantErr {
				t.Errorf("PlanetsUseCase.ListPlanets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			ok := assert.ElementsMatch(t, *got, *tt.want)

			if !ok {
				t.Errorf("PlanetsUseCase.ListPlanets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanetsUseCase_GetPlanetByName(t *testing.T) {
	type args struct {
		name string
	}
	type fields struct {
		Repo *repository.PlantesRespositoryMock
	}
	repo := new(repository.PlantesRespositoryMock)
	tests := []struct {
		name    string
		fields  fields
		want    *entity.Planet
		mock    *entity.Planet
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "GetPlanetByName",
			args: args{
				name: "Tatooine",
			},
			fields: fields{
				Repo: repo,
			},
			mock: &entity.Planet{
				Name: "Tatooine",
			},
			want: &entity.Planet{
				Name:         "Tatooine",
				Movies_count: 5,
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "GetPlanetByName-Error",
			args: args{
				name: "Tatooine",
			},
			fields: fields{
				Repo: repo,
			},
			mock:    nil,
			want:    nil,
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields.Repo.On("GetPlanetByName", "Tatooine").Return(tt.mock, tt.err).Once()

			u := &PlanetsUseCase{
				Repo: tt.fields.Repo,
			}
			got, err := u.GetPlanetByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlanetsUseCase.GetPlanetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlanetsUseCase.GetPlanetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanetsUseCase_GetPlanetById(t *testing.T) {
	type args struct {
		id string
	}
	type fields struct {
		Repo *repository.PlantesRespositoryMock
	}
	repo := new(repository.PlantesRespositoryMock)
	tests := []struct {
		name    string
		fields  fields
		want    *entity.Planet
		mock    *entity.Planet
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "GetPlanetById",
			args: args{
				id: "ID",
			},
			fields: fields{
				Repo: repo,
			},
			mock: &entity.Planet{
				Name: "Tatooine",
			},
			want: &entity.Planet{
				Name:         "Tatooine",
				Movies_count: 5,
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "GetPlanetById-Error",
			args: args{
				id: "ID",
			},
			fields: fields{
				Repo: repo,
			},
			mock:    nil,
			want:    nil,
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields.Repo.On("GetPlanetById", "ID").Return(tt.mock, tt.err).Once()

			u := &PlanetsUseCase{
				Repo: tt.fields.Repo,
			}
			got, err := u.GetPlanetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlanetsUseCase.GetPlanetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlanetsUseCase.GetPlanetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanetsUseCase_InsertPlanet(t *testing.T) {
	type fields struct {
		Repo *repository.PlantesRespositoryMock
	}
	type args struct {
		p *entity.Planet
	}
	repo := new(repository.PlantesRespositoryMock)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mongo.InsertOneResult
		err     error
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "InsertPlanet",
			args: args{
				p: &entity.Planet{Name: "Tatooine"},
			},
			fields: fields{
				Repo: repo,
			},
			want: &mongo.InsertOneResult{
				InsertedID: "ID",
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "InsertPlanet-Error",
			args: args{
				p: &entity.Planet{Name: "Tatooine"},
			},
			fields: fields{
				Repo: repo,
			},
			want:    nil,
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Repo.On("InsertPlanet", &entity.Planet{Name: "Tatooine"}).Return(tt.want, tt.err).Once()
			u := &PlanetsUseCase{
				Repo: tt.fields.Repo,
			}
			got, err := u.InsertPlanet(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlanetsUseCase.InsertPlanet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlanetsUseCase.InsertPlanet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanetsUseCase_DeletePlanetById(t *testing.T) {
	type args struct {
		id string
	}
	type fields struct {
		Repo *repository.PlantesRespositoryMock
	}
	repo := new(repository.PlantesRespositoryMock)
	tests := []struct {
		name    string
		fields  fields
		want    *mongo.DeleteResult
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "DeletePlanetById",
			args: args{
				id: "ID",
			},
			fields: fields{
				Repo: repo,
			},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "DeletePlanetById-Error",
			args: args{
				id: "ID",
			},
			fields: fields{
				Repo: repo,
			},
			want:    nil,
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Repo.On("DeletePlanetById", "ID").Return(tt.want, tt.err).Once()
			u := &PlanetsUseCase{
				Repo: tt.fields.Repo,
			}
			got, err := u.DeletePlanetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlanetsUseCase.DeletePlanetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlanetsUseCase.DeletePlanetById() = %v, want %v", got, tt.want)
			}
		})
	}
}
