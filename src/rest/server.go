package rest

import (
	"os"
	"path"

	"github.com/fredi12345/kuefa-karben/src/storage"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	db            storage.Service
	l             *zap.Logger
	imagePath     string
	thumbnailPath string
}

func NewServer(db storage.Service, logger *zap.Logger) *Server {
	logger = logger.Named("server")
	imagePath := path.Join(viper.GetString("web.storage"), "images")
	if err := os.MkdirAll(imagePath, 0750|os.ModeDir); err != nil {
		logger.Fatal("could not create folder", zap.Error(err))
	}

	thumbnailPath := path.Join(viper.GetString("web.storage"), "thumbnails")
	if err := os.MkdirAll(thumbnailPath, 0750|os.ModeDir); err != nil {
		logger.Fatal("could not create folder", zap.Error(err))
	}

	return &Server{
		db:            db,
		l:             logger,
		imagePath:     imagePath,
		thumbnailPath: thumbnailPath,
	}
}
