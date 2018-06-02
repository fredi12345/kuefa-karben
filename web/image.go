package web

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
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

	filename := s.writeFileToDisk(r)
	s.db.CreateImage("/public/images/"+filename, eventId)

	s.redirectToEventId(w, r, eventId)
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

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
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

	s.redirectToEventId(w, r, eventId)
}

func (s *Server) writeFileToDisk(r *http.Request) string {
	file, handler, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	infos, err := ioutil.ReadDir(s.imgPath)
	if err != nil {
		panic(err)
	}

	filename := strconv.Itoa(len(infos)) + handler.Filename

	f, err := os.Create(path.Join(s.imgPath, filename))
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, file)
	if err != nil {
		panic(err)
	}

	return filename
}
