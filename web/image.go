package web

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
)

func (s *Server) AddImage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		return err
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return err
	}

	filename := getUniqueFileName()
	file, _, err := r.FormFile("image")
	if err != nil {
		return err
	}
	defer file.Close()

	err = s.saveNewFile(file, filename)
	if err != nil {
		return err
	}

	err = s.db.CreateImage("/public/images/"+filename, eventId)
	if err != nil {
		return err
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)

	return nil
}

func (s *Server) DeleteImage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	imageId, err := strconv.Atoi(r.Form.Get("imageId"))
	if err != nil {
		return err
	}

	filename, err := s.db.DeleteImage(imageId)
	if err != nil {
		return err
	}

	err = os.Remove(path.Join(s.imgPath, filename))
	if err != nil {
		return err
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)

	return nil
}

func (s *Server) saveNewFile(file io.Reader, name string) error {
	f, err := os.Create(filepath.Join(s.imgPath, name))
	if err != nil {
		return err
	}

	_, err = io.Copy(f, file)
	return err
}

func getUniqueFileName() string {
	return strconv.Itoa(int(time.Now().UnixNano()))
}
