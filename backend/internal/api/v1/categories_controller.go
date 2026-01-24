package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web3crowdfunding/internal/repositories"
)

func StartCategoriesController(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/categories", getAllCategories)
}

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := repositories.GetAllCategories()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(categories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}
