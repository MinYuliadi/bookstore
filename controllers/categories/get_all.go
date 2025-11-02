package categories_controllers

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var categories []models.Category

	query := `
		SELECT id, name, created_at, created_by, modified_at, modified_by
		FROM categories
	`

	rows, errorRows := config.DB.Query(query)

	if errorRows != nil {
		utils.Error(c, http.StatusInternalServerError, errorRows.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var row models.Category
		if err := rows.Scan(&row.ID, &row.Name, &row.CreatedAt, &row.CreatedBy, &row.ModifiedAt, &row.ModifiedBy); err != nil {
			utils.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		categories = append(categories, row)
	}

	utils.Success(c, "Success", categories)
}
