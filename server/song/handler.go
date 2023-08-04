package song

import (
	"net/http"

	"music_platform/server/handlers"
	"music_platform/server/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

const (
	songsURL   = "/songs"
	songURL    = "/songs/:UUID"
	searchURL  = "/search"
)

type songHandler struct {
	logger *logging.Logger
}

func NewSongHandler(logger *logging.Logger) handlers.Handler {
	return &songHandler{
		logger: logger,
	}
}

func (h *songHandler) Register(router *httprouter.Router) {
	router.GET(songsURL, h.GetSongList)
	router.POST(songsURL, h.AddSong)
	router.GET(songURL, h.GetSongByUUID)
	router.PUT(songURL, h.UpdateSong)
	router.DELETE(songURL, h.DeleteSong)
	router.GET(searchURL, h.SearchSong)
}

func (h *songHandler) GetSongList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is the list of songs"))
}

func (h *songHandler) AddSong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is adding a new song"))
}

func (h *songHandler) GetSongByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is getting a song by UUID"))
}

func (h *songHandler) UpdateSong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is updating a song"))
}

func (h *songHandler) DeleteSong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is deleting a song"))
}

func (h *songHandler) SearchSong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	query := r.URL.Query().Get("query")
	w.Write([]byte("Searching for songs with query: " + query))
}
