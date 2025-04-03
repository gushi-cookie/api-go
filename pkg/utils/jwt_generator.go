package utils

import (
	"apigo/pkg/configs"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	Access           string
	Refresh          string
	AccessExpiresIn  time.Duration
	RefreshExpiresIn time.Duration
}

func GenerateNewTokens(id string) (*Tokens, error) {
	accessToken, accessExpIn, err := generateAccessToken(id)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpIn, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		Access:           accessToken,
		Refresh:          refreshToken,
		AccessExpiresIn:  *accessExpIn,
		RefreshExpiresIn: *refreshExpIn,
	}, nil
}

func generateAccessToken(id string) (string, *time.Duration, error) {
	config, err := configs.GetJWTConfig()
	if err != nil {
		return "", nil, err
	}

	expiresAt := time.Now().Add(time.Minute * time.Duration(config.ExpiresInMinutes))
	expiresIn := expiresAt.Sub(time.Now())

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = expiresAt.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.SecretKey))

	return signed, &expiresIn, err
}

func generateRefreshToken() (string, *time.Duration, error) {
	config, err := configs.GetJWTConfig()
	if err != nil {
		return "", nil, err
	}

	hash := sha256.New()
	content := strconv.Itoa(rand.Int()) + time.Now().String()
	_, err = hash.Write([]byte(content))
	if err != nil {
		return "", nil, err
	}

	expiresAt := time.Now().Add(time.Minute + time.Duration(config.RefreshKeyExpiresInMinutes))
	expiresIn := expiresAt.Sub(time.Now())

	token := hex.EncodeToString(hash.Sum(nil)) + "." + fmt.Sprint(expiresAt.Unix())

	return token, &expiresIn, nil
}
