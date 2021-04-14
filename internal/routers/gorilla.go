package gorilla


import (
	"github.com/gorilla/mux"
	"adverto/internal/api"
)




func RegisterRouter() *mux.Router {
	return mux.NewRouter()
}

func RegisterHandlers(r *mux.Router, a *api.Application) {
r.HandleFunc("/", a.AddProduct).Methods("POST")
r.HandleFunc("/allproducts", a.GetAllProducts).Methods("GET")
r.HandleFunc("/", a.FindOne).Methods("GET")
}