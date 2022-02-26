package rest

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"path"

	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type (
	// UploadImageRequest is the request definition for uploading images. It's possible to
	// upload a single image with multipart/form-data encoding.
	UploadImageRequest struct {
		// the image to upload
		// @Required
		Image *bytes.Buffer `json:"image" form:"image"`

		// will mark the image as title image for an event
		IsTitle bool `json:"isTitle" form:"isTitle"`
	}

	// UploadImageResponse is the response after an image was uploaded successfully.
	UploadImageResponse struct {
		// a UUIDv4 to identify the image
		// @Required
		ID string `json:"id"`

		// relative URL to access the image
		// @Required
		ImageURL string `json:"imageURL"`

		// relative URL to access a thumbnail of the image
		// @Required
		ThumbnailURL string `json:"thumbnailURL"`
	}
)

// UploadImage allows to upload an image.
//
// @OperationID UploadImage
// @Title upload an image
// @Param myimage form UploadImageRequest true "UploadImageRequest"
// @Success 200 object UploadImageResponse "Successfully uploaded the image"
// @Failure 400 object ErrorResponse "Error while uploading the image"
// @Failure 500 object ErrorResponse "Error while uploading the image"
// @Route /images [post]
func (s *Server) UploadImage(c echo.Context) error {
	// TODO return proper error structs

	var request UploadImageRequest
	err := c.Bind(&request)
	if err != nil {
		s.l.Error("could not bind request", zap.Error(err))
		return echo.ErrBadRequest
	}

	fullSizeImage, err := imaging.Decode(request.Image, imaging.AutoOrientation(true))
	if err != nil {
		s.l.Error("could not open form file", zap.Error(err))
		return echo.ErrInternalServerError
	}

	imageID, err := s.db.CreateTitleImage()
	if err != nil {
		s.l.Error("could not create title image", zap.Error(err))
		return echo.ErrInternalServerError
	}

	thumbnailImage := imaging.Fit(fullSizeImage, 400, 400, imaging.Lanczos)
	err = s.saveImage(thumbnailImage, path.Join(s.thumbnailPath, imageID))
	if err != nil {
		s.l.Error("could not save thumbnail", zap.Error(err))
		return echo.ErrInternalServerError
	}

	err = s.saveImage(fullSizeImage, path.Join(s.imagePath, imageID))
	if err != nil {
		s.l.Error("could not save image", zap.Error(err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, UploadImageResponse{
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

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return fmt.Errorf("could not encode image %s: %w", path, err)

	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("could not close file %s: %w", path, err)
	}

	return nil
}
