package api

import (
	"adverto/internal/models"
	"adverto/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

type Application struct {
	repo *repository.LinksRepository
}

func NewApplication(repo *repository.LinksRepository) *Application {
	return &Application{repo: repo}
}



func (a *Application) AddProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	log.Println(p)
	if err := a.repo.Create(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *Application) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	p, err := a.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	answer, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(answer)
}


func (a *Application) FindOne(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	product, err := a.repo.FindOne(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	answer, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(answer)
}
