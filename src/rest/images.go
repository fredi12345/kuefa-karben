package rest

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) UploadImage(c echo.Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Printf("could not find form file: %v", err)
		return echo.ErrBadRequest
	}

	f, err := fileHeader.Open()
	if err != nil {
		log.Printf("could not open form file: %v", err)
		return echo.ErrBadRequest
	}
	defer f.Close()

	var b bytes.Buffer
	_, err = io.Copy(&b, f)
	if err != nil {
		log.Printf("could not load file into buffer: %v", err)
		return echo.ErrInternalServerError
	}

	fullSizeImage, err := imaging.Decode(&b, imaging.AutoOrientation(true))
	if err != nil {
		log.Printf("could not open form file: %v", err)
		return echo.ErrInternalServerError
	}

	imageID := uuid.NewString()
	thumbnailImage := imaging.Fit(fullSizeImage, 400, 400, imaging.Lanczos)

	err = s.saveImage(thumbnailImage, path.Join(s.thumbnailPath, imageID))
	if err != nil {
		return err
	}

	err = s.saveImage(fullSizeImage, path.Join(s.imagePath, imageID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		ID           string `json:"id"`
		ImageURL     string `json:"imageURL"`
		ThumbnailURL string `json:"thumbnailURL"`
	}{
		ID:           imageID,
		ImageURL:     fmt.Sprintf("/public/images/%s.jpeg", imageID),
		ThumbnailURL: fmt.Sprintf("/public/thumbnails/%s.jpeg", imageID),
	})
}

func (s *Server) saveImage(img image.Image, path string) error {
	f, err := os.Create(path + ".jpeg")
	if err != nil {
		return fmt.Errorf("could not create file %s: %w", path, err)
	}
	defer f.Close()

	return jpeg.Encode(f, img, nil)
}
