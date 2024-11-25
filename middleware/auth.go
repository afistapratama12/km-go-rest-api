package middleware

import (
	"errors"
	"testrestapi/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		_, err := ValidateToken(authorization)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}

var (
	jwtKey = []byte("s3cr3tk3ydataINTERNAL")
)

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Buat claims berisi data username dan role yang akan kita embed ke JWT
	claims := &model.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// expiry time menggunakan time millisecond
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Buat JWT string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan (proses encoding JWT)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// return internal error ketika ada kesalahan saat pembuatan JWT string

		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tknStr string) (*model.Claims, error) {
	claims := &model.Claims{}

	// parse JWT token ke dalam claims
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// return unauthorized ketika ada kesalahan saat parsing token
			return nil, err
		}
		// return bad request ketika field token tidak ada
		return nil, err
	}

	//return unauthorized ketika token sudah tidak valid (biasanya karena token expired)
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
