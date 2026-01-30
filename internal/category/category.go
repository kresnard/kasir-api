package category

import (
	"cashier_api/entity"
	"cashier_api/helper"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var category = []entity.Category{
	{
		ID:          1,
		Name:        "Motor",
		Description: "Kendaraan roda dua",
	},
	{
		ID:          2,
		Name:        "Mobil",
		Description: "Kendaraan roda empat",
	},
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(helper.ContentType, helper.ApplicationJSON)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(
		helper.SuccessResponse("operasi berhasil", category),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, c := range category {
		if c.ID == id {
			w.Header().Set(helper.ContentType, helper.ApplicationJSON)
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(helper.SuccessResponse("operasi berhasil", c))
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var ctg entity.Category

	err := json.NewDecoder(r.Body).Decode(&ctg)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ctg.ID = len(category) + 1
	category = append(category, ctg)

	w.Header().Set(helper.ContentType, helper.ApplicationJSON)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(helper.SuccessResponse("operasi berhasil", nil))
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, c := range category {
		if c.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			category = append(category[:i], category[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "category belum ada", http.StatusNotFound)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	var updateCtg entity.Category
	err = json.NewDecoder(r.Body).Decode(&updateCtg)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, c := range category {
		if c.ID == id {
			updateCtg.ID = id
			category[i] = updateCtg

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(helper.SuccessResponse("operasi berhasil", nil))
			return
		}
	}

	http.Error(w, "category belum ada", http.StatusNotFound)
}
