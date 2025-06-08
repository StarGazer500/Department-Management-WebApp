package utils

import (
	"fmt"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Implement password hashing here
	// Example using bcrypt:
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func ComparePassword(requestPassword string, dbPassword string) bool {
	// Compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(requestPassword))
	if err != nil {
		// If there's an error (passwords don't match), return false and the error
		return false
	}
	// If no error (passwords match), return true and nil error
	return true
}


func LoadKeyAndReturnByte() ([]byte, error) {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		return nil, fmt.Errorf("missing JWT secret key in the environment")
	}
	return []byte(key), nil
}


func GenerateExpiryTime(durationInSeconds int64) int64 {
	expirationTime := time.Now().Add(time.Duration(durationInSeconds) * time.Second)
	return expirationTime.Unix()
}

func CheckTokenExpiration(claims jwt.MapClaims) error {
	if exp, ok := claims["exp"].(float64); ok {

		expirationTime := time.Unix(int64(exp), 0)

		if time.Now().After(expirationTime) {
			return fmt.Errorf("token has expired at %s", expirationTime)
		}

		fmt.Printf("Token is valid. Expiration time: %s\n", expirationTime)
		return nil
	} else {
		return fmt.Errorf("missing expiration claim in token")
	}
}


type TokenClaimStruct struct {
	MyAuthServer    string
	AuthUserEmail   string
	AuthUserId      int
	AuthExp         int64
}

func GenerateAccessToken(claim TokenClaimStruct) (string, error) {
	key, err := LoadKeyAndReturnByte()
	if err != nil {
		return "", fmt.Errorf("error loading the secret key: %w", err)
	}
	accessTokenExpiry := GenerateExpiryTime(30)
	claims := jwt.MapClaims{
		"iss":     claim.MyAuthServer,    // Issuer
		"sub":     claim.AuthUserEmail,   // Subject (user email)
	 // User surname
		"id":      claim.AuthUserId,      // User ID
		"exp":     accessTokenExpiry,     // for 120 seconds
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing the access token: %w", err)
	}

	return tokenString, nil
}

func GenerateRefreshToken(claim TokenClaimStruct) (string, error) {

	key, err := LoadKeyAndReturnByte()
	if err != nil {
		return "", fmt.Errorf("error loading the secret key: %w", err)
	}

	refreshTokenExpiry := GenerateExpiryTime(60) // 30 days * 24 hours * 60 minutes * 60 seconds

	claims := jwt.MapClaims{
		"iss":     claim.MyAuthServer,    // Issuer
		"sub":     claim.AuthUserEmail,   // Subject (user email)

		"id":      claim.AuthUserId,      // User ID
		"exp":     refreshTokenExpiry,    // Expiration time (Unix timestamp)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token and get the signed string
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing the refresh token: %w", err)
	}

	return tokenString, nil
}

func VerifyAccessToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing access token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if err := CheckTokenExpiration(claims); err != nil {
			fmt.Println("Access Expiry has happened")
			return nil, err
		}

		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid access token")
	}
}

func VerifyRefreshToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing refresh token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if err := CheckTokenExpiration(claims); err != nil {
			fmt.Println("Refresh Expiry has happened")
			return nil, err
		}

		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid refresh token")
	}
}