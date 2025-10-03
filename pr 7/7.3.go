// Корчагин Евгений 363
package main

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	Username string
	Email    string
	Password []byte
}

func (u *User) SetPassword(password string) {
	hash := sha256.New()
	hash.Write([]byte(password))
	u.Password = hash.Sum(nil)
}

func (u *User) VerifyPassword(password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	enteredPasswordHash := hash.Sum(nil)
	return string(u.Password) == string(enteredPasswordHash)
}

func main() {
	user := &User{
		Username: "lo",
		Email:    "lo@lololowka.ru",
	}

	user.SetPassword("12345678") //не взломать
	fmt.Println("Проверяем первый пароль", user.VerifyPassword("12345678"))
	fmt.Println("Проверяем второй пароль", user.VerifyPassword("44444444")) //  4 клубнички ,4 тарелки , 4 куска , 4 вилки   сплошные четверки

}
