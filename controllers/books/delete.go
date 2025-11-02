package books_controllers

import (
	"bookstore/config"
	"bookstore/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBook(c *gin.Context) {
	var id int
	idParam := c.Param("id")

	query := `
		DELETE
		FROM books
		WHERE id=$1
		RETURNING id
	`

	if err := config.DB.QueryRow(query, idParam).Scan(&id); err == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, fmt.Sprintf("data with id: %s not found", idParam))
		return
	} else if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Success", gin.H{
		"deletedId": id,
	})
}
