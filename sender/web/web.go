package web

type HttpServer struct {
	ip string
}

func New() *HttpServer {
	return &HttpServer{}
}
func (serv HttpServer) MustInit() {

}

func (serv HttpServer) Handle() error {
	return nil
}
