package auth

import (
	"homepage/pkg/configs"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// CreateToken create jwt
func CreateToken(studentID string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		configs.JWTStudentIDClaim: studentID,
		"exp":                     time.Now().Add(configs.JWTExpire).Unix(),
	}

	var secretKey = configs.JWTSecret
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		err = errors.Wrap(err, "failed to signed jwt")
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
		err = errors.Wrap(err, "failed to parse jwt")
		return token, err
	}

	return token, nil
}

// GetStudentIDFromJWT jwtのクレームから学籍番号の取得
func GetStudentIDFromJWT(token *jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	return claims[configs.JWTStudentIDClaim].(string)
}
