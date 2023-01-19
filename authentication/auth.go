package authentication

import (
	"context"
	"entdemo/ent"
	start "entdemo/start"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var input *ent.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	_, err := ent.EntClient.User.
		Create().
		SetUsername(input.Username).SetPassword(input.Password).Save(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "registered!"})
}
func Login(c *gin.Context) {

	var input *ent.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := LoginCheck(c, input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

func LoginCheck(c *gin.Context, username string, password string) (string, error) {

	user, err := start.QueryUserByName(c, ent.EntClient, username)
	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(uint(user.ID))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyPassword(password, s string) error {
	err := fmt.Errorf("password not matched")
	if password == s {
		return nil
	}
	return err
}
func CurrentUser(c *gin.Context) {

	user_id, err := ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := start.QueryUserById(context.Background(), ent.EntClient, int(user_id))
	fmt.Println("got current user", u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
