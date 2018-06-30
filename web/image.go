package web

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

func (s *Server) AddImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		panic(err)
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

	filename := getUniqueFileName()
	file, _, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s.saveNewFile(file, filename)
	s.db.CreateImage("/public/images/"+filename, eventId)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) DeleteImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	imageId, err := strconv.Atoi(r.Form.Get("imageId"))
	if err != nil {
		panic(err)
	}

	filename, err := s.db.DeleteImage(imageId)
	if err != nil {
		panic(err)
	}

	err = os.Remove(path.Join(s.imgPath, filename))
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
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
