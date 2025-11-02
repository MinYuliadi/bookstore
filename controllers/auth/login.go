package auth_controllers

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	query := `SELECT id, username, password FROM users WHERE username=$1`

	var user models.User

	if err := config.DB.QueryRow(query, req.Username).Scan(&user.ID, &user.Username, &user.Password); err == sql.ErrNoRows {
		utils.Error(c, http.StatusBadRequest, "username tidak terdaftar")
		return
	} else if err != nil && err != sql.ErrNoRows {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !utils.ComparePassword(req.Password, user.Password) {
		utils.Error(c, http.StatusBadRequest, "invalid password")
		return
	}

	token, errorJWT := utils.GenerateJWT(req.Username)

	if errorJWT != nil {
		utils.Error(c, http.StatusInternalServerError, errorJWT.Error())
		return
	}

	utils.Success(c, "Success", gin.H{
		"username": req.Username,
		"token":    token,
	})
}
