package books_controllers

import (
	"bookstore/config"
	"bookstore/helpers"
	"bookstore/models"
	"bookstore/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateBooks(c *gin.Context) {
	var book models.Book
	var id int
	var categoriesId []models.Category
	gotCategoryId := false
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
		INSERT INTO books
		(title, description, image_url, release_year, price, total_page, thickness, created_by, category_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	if err := config.DB.QueryRow(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, loggedInUser, book.CategoryID).Scan(&id); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Created(c, "Success", gin.H{
		"createdId": id,
	})
}
