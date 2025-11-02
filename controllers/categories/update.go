package categories_controllers

import (
	"bookstore/config"
	"bookstore/helpers"
	"bookstore/models"
	"bookstore/utils"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateCategories(c *gin.Context) {
	var category models.Category
	var reqCategory models.Category
	var existingName string
	id := c.Param("id")
	loggedInUser, isExist := c.Get(helpers.Username)

	if !isExist {
		utils.Error(c, http.StatusInternalServerError, "missing context: username")
		return
	}

	if err := c.ShouldBindBodyWithJSON(&reqCategory); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if reqCategory.Name == "" {
		utils.Error(c, http.StatusBadRequest, "nama tidak boleh kosong")
		return
	}

	getQuery := `
		SELECT name
		FROM categories
		WHERE name=$1
	`

	if err := config.DB.QueryRow(getQuery, reqCategory.Name).Scan(&existingName); err != nil && err != sql.ErrNoRows {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if strings.EqualFold(existingName, reqCategory.Name) {
		utils.Error(c, http.StatusBadRequest, "nama category sudah terdaftar")
		return
	}

	query := `
		UPDATE categories
		SET name=$1, modified_at=$2, modified_by=$3
		WHERE id=$4
		RETURNING id, name
	`

	if err := config.DB.QueryRow(query, reqCategory.Name, time.Now(), loggedInUser, id).Scan(&category.ID, &category.Name); err == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, fmt.Sprintf("data with id: %s not found", id))
		return
	} else if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success", gin.H{
		"updatedId": id,
		"name":      category.Name,
	})
}
