package auth

import (
	"time"

	"homepage/pkg/configs"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create jwt
func CreateToken(userID string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// TODO: claimsの設定
	token.Claims = jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = configs.JWTSecret
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

// VerifyToken validation jwt
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.JWTSecret), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}

// GetStudentIDFromJWT jwtのクレームから学籍番号の取得
func GetStudentIDFromJWT(token *jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	return claims["student_id"].(string)
}
