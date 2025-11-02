package auth_controllers

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.Error(c, http.StatusBadGateway, err.Error())
		return
	}

	password, err := utils.HashPassword(req.Password)

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	userQuery := `SELECT username FROM users WHERE username=$1`

	var username string

	if err := config.DB.QueryRow(userQuery, req.Username).Scan(&username); err != nil && err != sql.ErrNoRows {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if username == req.Username {
		utils.Error(c, http.StatusBadRequest, "username telah digunakan")
		return
	}

	query := `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id
	`

	var id int
	if err := config.DB.QueryRow(query, req.Username, password).Scan(&id); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Created(c, "Success", gin.H{
		"id":       id,
		"username": req.Username,
	})
}
