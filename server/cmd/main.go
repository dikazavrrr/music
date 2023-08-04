package main

import (
	"net"
	"net/http"
	"time"

	"music_platform/server/song"
	"music_platform/server/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("create user")
	handler := song.NewSongHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	logger.Info("server port 8080")
	logger.Fatalln(server.Serve(listener))
}