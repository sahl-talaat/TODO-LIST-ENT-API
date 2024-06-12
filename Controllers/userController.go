package controllers

import (
	"net/http"
	config "todoList/Config"
	"todoList/ent"
	"todoList/ent/user"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var input ent.User
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json data"})
		return
	}
	u, err := config.DB.User.Query().Where(user.UsernameEQ(input.Username)).Only(c)
	if err != nil {
		/* if ent.IsNotFound(err) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username"})
		} */
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))
	if input.Password != u.Password /* err != nil */ {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func SignUp(c *gin.Context) {
	var input ent.User
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json data"})
		return
	}
	u, err := config.DB.User.Create().SetUsername(input.Username).SetPassword(input.Password).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username is existing, must be uniqe" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}
