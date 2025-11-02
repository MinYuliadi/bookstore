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

func GetBooksByCategory(c *gin.Context) {
	paramId := c.Param("id")
	var books []models.Book

	query := `
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
		FROM books
		WHERE category_id=$1
	`

	bookRows, errGetQuery := config.DB.Query(query, paramId)

	if errGetQuery == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, fmt.Sprintf("data with id: %s not found", paramId))
		return
	} else if errGetQuery != nil {
		utils.Error(c, http.StatusInternalServerError, errGetQuery.Error())
		return
	}

	for bookRows.Next() {
		var book models.Book
		err := bookRows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)

		if err != nil {
			utils.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		books = append(books, book)
	}

	utils.Success(c, "Success", books)
}
