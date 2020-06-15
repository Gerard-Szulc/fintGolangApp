package users

import (
	"fintGolangApp/helpers"
	"fintGolangApp/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func prepareToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)
	return token
}

func Login(username string, pass string) map[string]interface{} {
	db := helpers.ConnectDB()
	user := &interfaces.User{}
	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}
	}

	accounts := []interfaces.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}
	defer db.Close()

	token := prepareToken(user)
	var response = map[string]interface{}{"message": "login_success"}
	response["jwt"] = token
	response["data"] = responseUser
	return response
}
