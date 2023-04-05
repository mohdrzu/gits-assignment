package models

import (
	"errors"
	"os"
	"time"

	"gits-assignment/config"

	"github.com/cristalhq/jwt/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
}

func (u *User) CreateUser(email, password string) error {
	var newUser User
	hashedPass, _ := newUser.HashPassword(password)

	newUser.Email = email
	newUser.Password = hashedPass

	err := config.DB.Create(&newUser).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) FindUserByEmail(email string) (User, error) {
	var user User
	result := config.DB.Where("email", email).Find(&user)
	if result.RowsAffected < 1 {
		return User{}, errors.New("record not found")
	}

	return user, nil
}

func (u *User) HashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (u *User) CheckPassword(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	return err == nil
}

func (u *User) GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 60)
	signer, err := jwt.NewSignerHS(jwt.HS256, []byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	claims := &jwt.RegisteredClaims{
		Issuer:    "gits-assignment",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Subject:   email,
	}
	builder := jwt.NewBuilder(signer)
	token, _ := builder.Build(claims)
	if token == nil {
		return "", errors.New("failed creating token")
	}

	return token.String(), nil
}
