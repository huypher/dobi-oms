package auth

import (
	"context"
	"errors"
	"time"

	"github.com/pghuy/dobi-oms/domain"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

func Login(ctx context.Context, username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("Invalid username or password")
	}

	//acc, err := s.userRepo.GetByUsername(ctx, username)
	//if err != nil {
	//	return "", err
	//}
	//if user == nil {
	//	return "", fmt.Errorf("User %s not exist", username)
	//}

	passErr := bcrypt.CompareHashAndPassword([]byte(""), []byte(password))
	if passErr != nil {
		return "", errors.New("Wrong password")
	}

	token, err := genJWT(&domain.Account{}, jwtTokenKey)
	if err != nil {
		logrus.WithError(err).Errorf("Gen JWT failed")
		return "", errors.New("Internal error")
	}

	return token, nil
}

func genJWT(acc *domain.Account, secretJWT string) (string, error) {
	if acc == nil {
		return "", errors.New("user empty")
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = acc.ID
	claims["username"] = acc.UserName
	claims["expired_time"] = time.Now().Add(time.Hour * 24).Unix()
	claims["token_type"] = "Bearer"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretJWT))
}

// use to hash user password when user create account
func genHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logrus.WithError(err).Errorf("Gen hash from pwd failed")
		return "", err
	}
	return string(hash), nil
}
