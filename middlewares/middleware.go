package middlewares

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"field-service/clients"
	"field-service/common/response"
	"field-service/config"
	"field-service/constants"
	errConstant "field-service/constants/error"

	"fmt"
	"net/http"
	"strings"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandlePanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				errMsg, _ := c.Get("error_message")
				httpStatus, _ := c.Get("http_status")
				message := fmt.Sprintf("%v", r)
				if errMsg != nil {
					message = fmt.Sprintf("%v", errMsg)
				}

				logrus.Errorf("Recovered from panic: %v", r)
				c.JSON(http.StatusInternalServerError, gin.H{
					"Code":    httpStatus,
					"Status":  errConstant.ErrInternalServerError.Error(),
					"Message": message,
					"Data":    nil,
					"Token":   nil,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func RateLimiter(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusTooManyRequests, response.Response{
				Status:  constants.Error,
				Message: errConstant.ErrTooManyRequests.Error(),
			})
			c.Abort()
		}
		c.Next()
	}
}

func extractBearerToken(token string) string {
	arrayToken := strings.Split(token, " ")
	if len(arrayToken) == 2 {
		return arrayToken[1]
	}
	return ""
}

func responseUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, response.Response{
		Status: constants.Error, Message: message,
	})
	c.Abort()
}

func validateApiKey(c *gin.Context) error {
	apiKey := c.GetHeader(constants.XApiKey)
	requestAt := c.GetHeader(constants.XRequestAt)
	serviceName := c.GetHeader(constants.XServiceName)
	signatureKey := config.Config.SignatureKey

	validateKey := fmt.Sprintf("%s:%s:%s", serviceName, signatureKey, requestAt)
	hash := sha256.New()
	hash.Write([]byte(validateKey))
	resultHash := hex.EncodeToString(hash.Sum(nil))

	if apiKey != resultHash {
		return errConstant.ErrUnauthorized
	}
	return nil
}

// func validateBearerToken(c *gin.Context, token string) error {
// 	if !strings.Contains(token, "Bearer") {
// 		fmt.Println("tes validateBearerToken kentoken ", token)
// 		return errConstant.ErrUnauthorized
// 	}
// 	tokenString := extractBearerToken(token)
// 	fmt.Println("tes validateBearerToken extractBearerToken ", tokenString)

// 	if tokenString == "" {
// 		return errConstant.ErrUnauthorized
// 	}
// 	claims := &userServices.Claims{}
// 	fmt.Println("tes validateBearerToken claims ", claims)

// 	tokenJwt, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)
// 		fmt.Println("tes validateBearerToken ok ", ok)

// 		if !ok {
// 			return nil, errConstant.ErrInvalidToken
// 		}
// 		jwtSecret := []byte(config.Config.JwtSecretKey)
// 		fmt.Println("tes validateBearerToken jwtSecret ", jwtSecret)

// 		return jwtSecret, nil
// 	})
// 	fmt.Println("tes validateBearerToken tokenJwt ", tokenJwt, "err: ", err)
// 	fmt.Println("tes validateBearerToken tokenJwt ", tokenJwt.Valid, "err: ", err)

// 	if err != nil || !tokenJwt.Valid {
// 		return errConstant.ErrUnauthorized
// 	}
// 	userLogin := c.Request.WithContext(context.WithValue(c.Request.Context(), constants.UserLogin, claims.User))
// 	c.Request = userLogin
// 	c.Set(constants.Token, token)
// 	return nil
// }

func contains(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func CheckRole(roles []string, clients clients.IClientRegistry) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			fmt.Println("tes CheckRole: missing Authorization header")
			responseUnauthorized(c, errConstant.ErrUnauthorized.Error())
			return
		}

		// Inject token into context
		ctx := context.WithValue(c.Request.Context(), constants.Token, token)

		user, err := clients.GetUser().GetUserByToken(ctx)
		if err != nil {
			fmt.Println("tes CheckRole: error getting user:", err)
			responseUnauthorized(c, errConstant.ErrUnauthorized.Error())
			return
		}

		if !contains(roles, user.Role) {
			fmt.Println("tes CheckRole: ga ada role ", roles, " user.Role ", user.Role)
			responseUnauthorized(c, errConstant.ErrUnauthorized.Error())
			return
		}

		c.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var err error
		token := c.GetHeader(constants.Authorization)
		fmt.Println("tes Authenticate kentoken ", token)
		if token == "" {
			fmt.Println("tes Authenticate token koshong ", token)
			responseUnauthorized(c, errConstant.ErrUnauthorized.Error())
			return
		}
		// err = validateBearerToken(c, token)
		// fmt.Println("tes Authenticate validateBearerToken ", err)

		// if err != nil {
		// 	responseUnauthorized(c, err.Error())
		// 	return
		// }

		// err = validateApiKey(c)
		// fmt.Println("tes Authenticate validateApiKey ", err)

		// if err != nil {
		// 	responseUnauthorized(c, err.Error())
		// 	return
		// }
		c.Next()
	}
}

func AuthenticateWithoutToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
