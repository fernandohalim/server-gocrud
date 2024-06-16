package productcontroller

import (
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
