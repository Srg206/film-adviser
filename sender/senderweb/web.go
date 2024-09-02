package senderweb

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
}

type FilmToSave struct {
	Name string `json: "name"`
}

func New() *HttpServer {
	return &HttpServer{}
}
func (serv *HttpServer) MustInit() {

}

func (serv HttpServer) Handle() error {

	router := gin.Default()

	router.POST("/movie", func(c *gin.Context) {
		var film FilmToSave
		// Декодирование JSON из тела запроса
		if err := c.BindJSON(&film); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Сохранение данных (в данном случае просто вывод в консоль)
		fmt.Printf("Получены данные: %+v\n", film)
		c.JSON(http.StatusOK, gin.H{"message": "Данные получены успешно"})
	})
	router.Run(":8000")
	return nil
}
