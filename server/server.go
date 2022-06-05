package server

import "groundswell.io/datastore"

type Server struct {
	DB *datastore.Map
	// Here goes all the info for a server (port, host, etc) in order to connect with several clients
}

var Ser *Server

func Init() {
	Ser = &Server{
		DB: datastore.NewMap(),
	}
}
