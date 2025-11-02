package books_controllers

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var books []models.Book

	query := `
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
		FROM books
	`

	bookRows, getErr := config.DB.Query(query)

	if getErr != nil {
		utils.Error(c, http.StatusInternalServerError, getErr.Error())
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
