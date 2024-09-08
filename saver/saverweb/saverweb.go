package saverweb

import (
	"film-adviser/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpSaver struct {
	repo repository.Repository
}

type FilmToSave struct {
	Userid int64  `json: "userid"`
	Name   string `json: "name"`
}

func New() HttpSaver {
	return HttpSaver{}
}
func (serv HttpSaver) MustInit(repo repository.Repository) {
	serv.repo = repo
}

func (serv HttpSaver) Handle() error {

	router := gin.Default()

	router.POST("/movie", func(c *gin.Context) {
		var film FilmToSave
		// Decode json from request body
		if err := c.BindJSON(&film); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received data: %+v\n", film)
		serv.repo.Write(film.Userid, film.Name)
		c.JSON(http.StatusOK, gin.H{"message": "Data received successully"})
	})
	router.Run(":8000")
	return nil
}
