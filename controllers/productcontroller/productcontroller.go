package productcontroller

import (
	"encoding/json"
	"net/http"
	"server-hydraulicpump/config"
	"server-hydraulicpump/helper"
	"server-hydraulicpump/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var product []models.Product

	if err := config.DB.Find(&product).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Product List", product)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&product).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Product Data Successfully Created", nil)
}

// func Detail(w http.ResponseWriter, r *http.Request) {
// 	idParams := mux.Vars(r)["id"]
// 	id, _ := strconv.Atoi(idParams)

// 	var product models.Product

// 	if err := config.DB.First(&product, id).Error; err!= {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			helper.Response(w, 404, "Product not Found", nil)
// 			return
// 		}
// 	}

// }
