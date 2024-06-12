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
	err = setLogIn(c, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't log in " + err.Error()})
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
	u, err := config.DB.User.Create().SetUsername(input.Username).SetPassword(input.Password).SetLogedin(true).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username is existing, must be uniqe ..." + err.Error()})
		return
	}
	err = setLogIn(c, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't log in " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}

func LogOut(c *gin.Context) {
	username := c.Param("username")
	u, err := checkUsername(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username in header not found" + err.Error()})
		return
	}
	err = setLogOut(c, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't log out " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "loged out successffuly"})
}

func UpdateUser(c *gin.Context) {
	username := c.Param("username")
	var input ent.User
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json data"})
		return
	}
	u, err := checkUsername(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username in header not found" + err.Error()})
		return
	}
	/* _, err = checkUsername(c, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username in pody is already existing" + err.Error()})
		return
	} */
	err = config.DB.User.UpdateOne(u).SetUsername(input.Username).SetPassword(input.Password).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, u)
}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")
	u, err := checkUsername(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username in header not found" + err.Error()})
		return
	}
	err = config.DB.User.DeleteOne(u).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successffuly"})
}

func checkUsername(c *gin.Context, username string) (u *ent.User, err error) {
	u, err = config.DB.User.Query().Where(user.UsernameEQ(username)).Only(c)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func setLogIn(c *gin.Context, u *ent.User) error {
	u, err := config.DB.User.UpdateOne(u).SetLogedin(true).Save(c)
	if err != nil {
		return err
	}
	return nil
}

func setLogOut(c *gin.Context, u *ent.User) error {
	u, err := config.DB.User.UpdateOne(u).SetLogedin(false).Save(c)
	if err != nil {
		return err
	}
	return nil
}
