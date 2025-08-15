package auth

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go-boilerplate/config"
	"go-boilerplate/modules/auth/dto"
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/models"
	"golang.org/x/crypto/pbkdf2"
	"time"
)

func Register(registerDTO dto.RegisterDTO) (error, models.User) {
	var user models.User

	user.Name = registerDTO.Name
	user.Email = registerDTO.Email
	user.Salt = uuid.New().String()

	passwordHashed := pbkdf2.Key([]byte(registerDTO.Password), []byte(user.Salt), 4096, 100, sha512.New)

	user.Password = hex.EncodeToString(passwordHashed)

	if result := core.DB.Create(&user); result.Error != nil {
		return result.Error, user
	}

	return nil, user
}

func Login(loginDTO dto.LoginDTO) (err error, loginResponseDTO dto.LoginResponseDTO) {
	var user models.User

	if result := core.DB.Where("email = ?", loginDTO.Email).Take(&user); result.Error != nil {
		err = result.Error
		return
	}

	if user.ID == uuid.Nil {
		err = errors.New("username_or_password_failed")
		return
	}

	err, accessToken, refreshToken := GenerateToken(user.ID.String())

	loginResponseDTO.AccessToken = accessToken
	loginResponseDTO.RefreshToken = refreshToken

	return
}

func GenerateToken(userId string) (err error, accessToken string, refreshToken string) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err = token.SignedString([]byte(config.App.JwtSecret))

	if err != nil {
		return
	}

	refreshToken = uuid.New().String()

	return
}
