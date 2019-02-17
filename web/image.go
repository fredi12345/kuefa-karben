package web

import (
	"bytes"
	"fmt"
	"github.com/fredi12345/kuefa-karben/web/template"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
			if err.Error() == "image: unknown format" {
				sess.AddFlash(&template.Message{Type: template.TypeError, Text: "Hinzufügen von Bildern fehlgeschlagen. Mindestens eine Datei hatte ein unbekanntes Format. Unterstützte Bildformate sind .jpg und .png."})
				_ = sess.Save(r, w)
				http.Redirect(w, r, fmt.Sprintf("/event/%d", eventId), http.StatusSeeOther)
			} else {
				return err
			}
		}

		err = s.db.CreateImage(filename, eventId)
		if err != nil {
			return errors.WithMessage(err, "cannot create image")
		}
	}

	http.Redirect(w, r, fmt.Sprint(r.Referer()+"#images"), http.StatusSeeOther)
	return nil
}

func (s *Server) createAndSaveThumbAndFullImage(filename string, file io.Reader) error {
	buffer, err := ioutil.ReadAll(file)

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

	err = s.deleteImageById(imageId)
	if err != nil {
		return errors.WithMessage(err, "cannot remove image "+strconv.Itoa(imageId))
	}

	http.Redirect(w, r, fmt.Sprint(r.Referer()+"#images"), http.StatusSeeOther)
	return nil
}

func (s *Server) deleteImageById(id int) error {
	filename, err := s.db.DeleteImage(id)
	if err != nil {
		return errors.WithMessage(err, "cannot delete image "+strconv.Itoa(id))
	}

	err = s.removeImageFileByFilename(filename)
	if err != nil {
		return errors.WithMessage(err, "cannot remove image "+strconv.Itoa(id))
	}

	return nil
}

func (s *Server) removeImageFileByFilename(filename string) error {
	err := os.Remove(filepath.Join(s.thumbPath, filename))
	if err != nil && !os.IsNotExist(err) {
		return errors.WithStack(err)
	}

	err = os.Remove(filepath.Join(s.imgPath, filename))
	if err != nil && !os.IsNotExist(err) {
		return errors.WithStack(err)
	}
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
