package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func jwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken, err := extractBearerToken(ctx.GetHeader("Authorization"))
		if err != nil {
			serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
				"data": map[string]interface{}{
					"success": false,
					"message": "token receive failed",
					"error":   err.Error(),
				},
			}, err)
			ctx.Abort()
			return
		}

		fmt.Println(jwtToken)

		token, err := verifyJwtToken(jwtToken)
		if err != nil {
			serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
				"data": map[string]interface{}{
					"success": false,
					"message": "verify token failed",
					"error":   err.Error(),
				},
			}, err)
			ctx.Abort()
			return
		}

		claims, OK := token.Claims.(jwt.MapClaims)
		if !OK {
			serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
				"data": map[string]interface{}{
					"success": false,
					"message": "Operation failed",
					"error":   fmt.Errorf("authentication error happend").Error(),
				},
			}, fmt.Errorf("authentication error happend"))
			ctx.Abort()
			return
		}

		username, OK := claims["username"].(string)
		if !OK {
			serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
				"data": map[string]interface{}{
					"success": false,
					"message": "Operation failed",
					"error":   fmt.Errorf("authentication error happend").Error(),
				},
			}, fmt.Errorf("authentication error happend"))
			ctx.Abort()
			return
		}

		// ---------------------------------specific routes---------------------------------
		// url := ctx.Request.URL
		// if strings.HasPrefix(url.String(),"/api/v1/posts") {
		// 	ctx.Set("username", username)
		// }
		// ---------------------------------------------------------------------------------
		ctx.Set("username", username)

		ctx.Next()
	}
}

func createJwtToken(username string) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 200).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyJwtToken(tokenString string) (*jwt.Token, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func getSecretKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	secretKey := os.Getenv("JWT_SECRET_TOKEN")
	return secretKey, nil
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", fmt.Errorf("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", fmt.Errorf("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
