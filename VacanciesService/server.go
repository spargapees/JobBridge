package server

import "net/http"

type HTTPServer struct {
	My_server *http.Server
}

func (srv *HTTPServer) Run_server(port string, handler http.Handler) error {
	srv.My_server = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return srv.My_server.ListenAndServe()
}
