package server

type HTTP3 interface {
	StartHTTP3()
}

type HTTP1 interface {
	StartHTTP1()
}

type Server struct {
}

func (s *Server) StartHTTP3() {
	startHTTP3Server()
}

func (s *Server) StartHTTP1() {
	startHTTP1Server()
}
