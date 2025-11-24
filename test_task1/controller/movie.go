package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github/muhammadyusuf/http/models"
	"github/muhammadyusuf/http/storage"
)

func (c *Controller) Movie(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		c.CreateMovie(w, r)
	}
	if r.Method == "GET" {
		path := strings.Split(r.URL.Path, "/")
		fmt.Println("path:", path)
		if len(path) > 2 {
			c.GetByIdMovie(w, r)
		} else {
			c.GetAllMovie(w, r)
		}
	}
	if r.Method == "PUT" {

		c.UpdateMovie(w, r)
	}

	if r.Method == "DELETE" {
		c.DeleteMovie(w, r)
	}

}

func (c *Controller) CreateMovie(w http.ResponseWriter, r *http.Request) {

	var movie models.Movie

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := storage.InsertMovie(c.db, movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	movie, err = storage.GetByIdMovie(c.db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (c *Controller) GetByIdMovie(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/movie/")

	movie, err := storage.GetByIdMovie(c.db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) GetAllMovie(w http.ResponseWriter, r *http.Request) {

	movie, err := storage.GetAllMovie(c.db)
	if err != nil {
		log.Println("3rd", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		log.Println("4th", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("1st", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = storage.UpdateMovie(c.db, movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode("succesfully updated")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/movie/")

	err := storage.DeleteMovie(c.db, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	err = json.NewEncoder(w).Encode("succesfully deleted!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
}
