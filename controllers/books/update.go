package books_controllers

import (
	"bookstore/config"
	"bookstore/helpers"
	"bookstore/models"
	"bookstore/utils"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateBook(c *gin.Context) {
	var book models.Book
	var id int
	var categoriesId []models.Category
	gotCategoryId := false
	idParam := c.Param("id")
	loggedInUser, isExist := c.Get(helpers.Username)

	if !isExist {
		utils.Error(c, http.StatusInternalServerError, "missing context: username")
		return
	}

	if err := c.ShouldBindBodyWithJSON(&book); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if book.TotalPage >= 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		utils.Error(c, http.StatusBadRequest, "tahun release hanya boleh diisi dengan range 1980 sampai 2024")
		return
	}

	categoryQuery := `
		SELECT id
		FROM categories
	`

	categoriesRows, errGetCategory := config.DB.Query(categoryQuery)

	if errGetCategory != nil {
		utils.Error(c, http.StatusInternalServerError, errGetCategory.Error())
		return
	}

	for categoriesRows.Next() {
		var row models.Category

		if err := categoriesRows.Scan(&row.ID); err != nil {
			utils.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		categoriesId = append(categoriesId, row)
	}

	for _, item := range categoriesId {
		if strings.Contains(strconv.Itoa(item.ID), strconv.Itoa(book.CategoryID)) {
			gotCategoryId = true
			break
		}
	}

	if !gotCategoryId {
		utils.Error(c, http.StatusBadRequest, fmt.Sprintf("category id: %d belum terdaftar pada table category", book.CategoryID))
		return
	}

	query := `
		UPDATE books
		SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, category_id=$8, modified_at=$9, modified_by=$10
		WHERE id=$11
		RETURNING id
	`

	err := config.DB.QueryRow(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, time.Now(), loggedInUser, idParam).Scan(&id)

	if err == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, fmt.Sprintf("data with id: %s not found", idParam))
		return
	}

	utils.Success(c, "Success", gin.H{
		"updatedId": id,
	})
}
