package receiverweb

type HttpServer struct {
}

func (serv *HttpServer) MustInit() {

}

func (serv HttpServer) PickFilm(chatid int64) string {

	return "Pulp fiction"
}

func (serv HttpServer) SendAnswer() {

}
