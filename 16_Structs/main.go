package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName       string
	lastName        string
	age             int
	description     string
	avatar          string
	primaryLanguage string
	createdAt       time.Time
}

func (u *user) chnagePrimaryLanguage(lang string) {
	u.primaryLanguage = lang
}

func (u user) getAllDetails() {
	user := []any{
		u.firstName,
		u.lastName,
		u.age,
		u.avatar,
		u.description,
		u.primaryLanguage,
	}
	for _, v := range user {
		fmt.Println("->", v)
	}
}

func main() {
	user1 := user{
		firstName:       "Ayush",
		lastName:        "Choudhary",
		age:             19,
		description:     "Full stack web app Develoepr",
		avatar:          "image",
		primaryLanguage: "JavaScript",
	}
	user1.createdAt = time.Now()
	fmt.Println(user1)
	fmt.Println(user1.primaryLanguage)
	user1.chnagePrimaryLanguage("GoLang")
	fmt.Println(user1.primaryLanguage)

	user1.getAllDetails()

}
