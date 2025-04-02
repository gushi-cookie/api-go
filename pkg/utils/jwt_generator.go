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
	Access  string
	Refresh string
}

func GenerateNewTokens(id string) (*Tokens, error) {
	accessToken, err := generateAccessToken(id)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateAccessToken(id string) (string, error) {
	config, err := configs.GetJWTConfig()
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.ExpiresInMinutes)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.SecretKey))
}

func generateRefreshToken() (string, error) {
	config, err := configs.GetJWTConfig()
	if err != nil {
		return "", err
	}

	hash := sha256.New()
	content := strconv.Itoa(rand.Int()) + time.Now().String()
	_, err = hash.Write([]byte(content))
	if err != nil {
		return "", err
	}

	expiresIn := time.Now().Add(time.Minute + time.Duration(config.RefreshKeyExpiresInMinutes)).Unix()

	token := hex.EncodeToString(hash.Sum(nil)) + "." + fmt.Sprint(expiresIn)

	return token, nil
}
