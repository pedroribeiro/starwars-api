package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroribeiro/starwars-api/internal/domain/entity"
	"github.com/pedroribeiro/starwars-api/internal/usecase"
)

type PlanetHandler struct {
	PlanetsUseCase usecase.IPlanetsUseCase
}

func (h *PlanetHandler) InitPlanetRoutes(router *gin.Engine) {
	router.GET("/planets", h.SearchPlanets)
	router.POST("/planets", h.InserPlanet)
	router.DELETE("/planets", h.DeletePlanetById)
}

func (h *PlanetHandler) InserPlanet(c *gin.Context) {
	var p entity.Planet

	if err := c.BindJSON(&p); err != nil {
		c.JSON(406, gin.H{"error": err.Error(), "form": p})
		c.Abort()
		return
	}

	res, err := h.PlanetsUseCase.InsertPlanet(&p)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, res)
	}
}

func (h *PlanetHandler) SearchPlanets(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")

	if name != "" {
		p, err := h.PlanetsUseCase.GetPlanetByName(name)

		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				c.JSON(404, gin.H{"message": "no results found"})
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(200, p)
		return
	}

	if id != "" {
		p, err := h.PlanetsUseCase.GetPlanetById(id)

		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				c.JSON(404, gin.H{"message": "no results found"})
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(200, p)
		return
	}

	p, err := h.PlanetsUseCase.ListPlanets()

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, p)
}

func (h *PlanetHandler) DeletePlanetById(c *gin.Context) {
	id := c.Query("id")

	if id != "" {
		p, err := h.PlanetsUseCase.DeletePlanetById(id)

		if err != nil {
			c.JSON(500, err)
			return
		}

		c.JSON(200, p)
		return
	}

	c.JSON(400, "Bad Request")
}
