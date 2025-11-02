package categories_controllers

import (
	"bookstore/config"
	"bookstore/helpers"
	"bookstore/models"
	"bookstore/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategories(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindBodyWithJSON(&category); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	getCategoryQuery := `SELECT name FROM categories WHERE name=$1`

	var categoryName string

	if err := config.DB.QueryRow(getCategoryQuery, category.Name).Scan(&categoryName); err != sql.ErrNoRows && err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if categoryName == category.Name {
		utils.Error(c, http.StatusBadRequest, "kategori sudah terdaftar")
		return
	}

	loggedInUser, isExist := c.Get(helpers.Username)

	if !isExist {
		utils.Error(c, http.StatusInternalServerError, "missing context: username")
		return
	}

	query := `
		INSERT INTO categories
		(name, created_by)
		VALUES($1, $2)
		RETURNING id
	`

	var id int
	if err := config.DB.QueryRow(query, category.Name, loggedInUser).Scan(&id); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Created(c, "Success", gin.H{
		"id":       id,
		"category": category.Name,
	})
}
