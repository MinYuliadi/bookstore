package categories_controllers

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoriesById(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	query := `
		SELECT id, name, created_at, created_by, modified_at, modified_by
		FROM categories
		WHERE id=$1
	`

	if err := config.DB.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy); err == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, fmt.Sprintf("data with id: %s not found", id))
		return
	} else if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success", category)
}
