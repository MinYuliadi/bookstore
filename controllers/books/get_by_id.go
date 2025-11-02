package books_controllers

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookById(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	query := `
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
		FROM books
		WHERE id=$1
	`

	errGet := config.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)

	if errGet == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, fmt.Sprintf("data with id: %s not found", id))
		return
	} else if errGet != nil {
		utils.Error(c, http.StatusInternalServerError, errGet.Error())
		return
	}

	utils.Success(c, "Success", book)
}
