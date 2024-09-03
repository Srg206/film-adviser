package receiverweb

import (
	"film-adviser/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	repo *repository.Repository
}

type FilmToRecomend struct {
	Name string `json: "name"`
}

//Methods

func New() *HttpServer {
	return &HttpServer{}
}

func (serv *HttpServer) MustInit() {

}

func (serv HttpServer) PickFilm(repo *repository.Repository, chatid int64) string {

	return "Pulp fiction"
}

func (serv HttpServer) SendAnswer() {
	router := gin.Default()
	router.GET("/movie", func(c *gin.Context) {
		chatID := c.Query("chatid")
		id, _ := strconv.ParseInt(chatID, 10, 64)
		response := FilmToRecomend{Name: serv.PickFilm(serv.repo, id)}
		c.JSON(http.StatusOK, response)
	})
	router.Run(":8080")
}
