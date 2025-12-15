import (
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.HandleFunc("/me", controllers.Me).Methods("GET")
}