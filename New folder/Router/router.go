package router

import(
	"postgress/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/stocks/{id}", middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/newstock", middleware.GetStock).Methods("POST","OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deletestock/{id}", middleware.GetStock).Methods("DELETE","OPTIONS")

}