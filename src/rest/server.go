package rest

import (
	"log"
	"os"
	"path"

	"github.com/fredi12345/kuefa-karben/src/storage"
	"github.com/spf13/viper"
)

type Server struct {
	db            storage.Service
	imagePath     string
	thumbnailPath string
}

func NewServer(db storage.Service) *Server {
	imagePath := path.Join(viper.GetString("web.storage"), "images")
	if err := os.MkdirAll(imagePath, 0750|os.ModeDir); err != nil {
		log.Fatalf("could not create folder: %v\n", err)
	}

	thumbnailPath := path.Join(viper.GetString("web.storage"), "thumbnails")
	if err := os.MkdirAll(thumbnailPath, 0750|os.ModeDir); err != nil {
		log.Fatalf("could not create folder: %v\n", err)
	}

	return &Server{
		db:            db,
		imagePath:     imagePath,
		thumbnailPath: thumbnailPath,
	}
}
