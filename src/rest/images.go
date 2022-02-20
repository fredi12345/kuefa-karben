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

type (
	// swagger:parameters UploadImage
	UploadImageRequest struct {
		// in: formData
		// required: true
		// swagger:file
		Image *bytes.Buffer `json:"image"`

		// in: formData
		// required: false
		IsTitle bool `json:"isTitle"`
	}

	UploadImageRequestData struct {
	}

	// swagger:response UploadImageResponse
	UploadImageResponse struct {
		// in: body
		Body UploadImageResponseData
	}

	UploadImageResponseData struct {
		// Generated v4 UUID of the image
		// Required: true
		// Example: 13cc859d-a679-49ff-9791-d62f3e761253
		ID string `json:"id"`

		// Relative URL to access the image
		// Required: true
		// Example: /public/image/13cc859d-a679-49ff-9791-d62f3e761253.jpeg
		ImageURL string `json:"imageURL"`

		// Relative URL to access the thumbnail
		// Required: true
		// Example: /public/thumbnails/13cc859d-a679-49ff-9791-d62f3e761253.jpeg
		ThumbnailURL string `json:"thumbnailURL"`
	}
)

// swagger:route POST /images kuefa UploadImage
//
// Upload an image.
//
// It is automatically served on /public/images/{id}. There is also a thumbnail being generated on /public/thumbnails/{id}.
//
// Consumes:
//   - multipart/form-data
//
// Responses:
//   200: UploadImageResponse
// 	 400: ErrorResponse
// 	 500: ErrorResponse
func (s *Server) UploadImage(c echo.Context) error {
	// TODO bind request?
	// TODO return proper error structs

	fileHeader, err := c.FormFile("image")
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

	return c.JSON(http.StatusOK, UploadImageResponseData{
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
