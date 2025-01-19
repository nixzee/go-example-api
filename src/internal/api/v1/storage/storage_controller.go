package storage

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type StorageController interface {
	RegisterRoutes() (router *chi.Mux)
	List(w http.ResponseWriter, r *http.Request)
}

var _ StorageController = (*storageController)(nil)

type storageController struct {
	storageService StorageService
}

func NewStorageController(storageService StorageService) StorageController {
	return &storageController{
		storageService: storageService,
	}
}

func (s *storageController) RegisterRoutes() (router *chi.Mux) {
	router = chi.NewRouter()
	router.Get("/", s.List)
	return
}

func (s *storageController) List(w http.ResponseWriter, r *http.Request) {
	// Get the container name form the query parameter
	containerName := chi.URLParam(r, "containerName")
	// Use the service to get the list of files (blobs) in the container
	blobs, err := s.storageService.ListFilesInContainer(r.Context(), containerName)
	if err != nil {
		http.Error(w, "Error listing files", http.StatusInternalServerError)
		return
	}
	// Set the headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Marshal the list of BlobInfo to JSON
	if err := json.NewEncoder(w).Encode(blobs); err != nil {
		http.Error(w, "Error encoding response to JSON", http.StatusInternalServerError)
	}
}
