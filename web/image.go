package web

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/nfnt/resize"

	"github.com/gorilla/sessions"
)

func (s *Server) AddImage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		return errors.WithStack(err)
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return errors.WithStack(err)
	}

	m := r.MultipartForm

	files := m.File["images"]

	for i := range files {

		filename := getUniqueFileName()
		file, err := files[i].Open()
		if err != nil {
			return errors.WithStack(err)
		}
		defer file.Close()

		err = s.createAndSaveThumbAndFullImage(filename, file)
		if err != nil {
			return err
		}

		err = s.db.CreateImage(filename, eventId)
		if err != nil {
			return errors.Wrap(err, "cannot create image")
		}
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	return nil
}

func (s *Server) createAndSaveThumbAndFullImage(filename string, file io.Reader) error {
	buffer, err := ioutil.ReadAll(file)

	//TODO: errorhandling wenn datei gar keine Bilddatei ist
	img, _, err := image.Decode(bytes.NewReader(buffer))
	if err != nil {
		return errors.WithStack(err)
	}

	thumbnailImage := resize.Thumbnail(400, 400, img, resize.Bilinear)

	err = s.saveNewThumbnailImage(thumbnailImage, filename)
	if err != nil {
		return err
	}

	err = s.saveNewFullImageFile(bytes.NewReader(buffer), filename)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) DeleteImage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	imageId, err := strconv.Atoi(r.Form.Get("imageId"))
	if err != nil {
		return errors.WithStack(err)
	}

	filename, err := s.db.DeleteImage(imageId)
	if err != nil {
		return errors.Wrap(err, "cannot delete image "+strconv.Itoa(imageId))
	}

	err = os.Remove(path.Join(s.imgPath, filename))
	if err != nil {
		return errors.WithStack(err)
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	return nil
}

func (s *Server) saveNewThumbnailImage(img image.Image, name string) error {
	f, err := os.Create(filepath.Join(s.thumbPath, name))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	return jpeg.Encode(f, img, nil)
}
func (s *Server) saveNewFullImageFile(img io.Reader, name string) error {
	f, err := os.Create(filepath.Join(s.imgPath, name))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	_, err = io.Copy(f, img)
	return errors.WithStack(err)
}

func getUniqueFileName() string {
	return strconv.Itoa(int(time.Now().UnixNano())) + ".jpg"
}
