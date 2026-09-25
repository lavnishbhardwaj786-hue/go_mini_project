package main

import (
	"errors"
	"fmt"
	"os"
)

type usererror struct {
	name string
}

type passerror struct {
	password string
}

func (u usererror) Error() string {
	return "invalid user name " + u.name
}
func (u passerror) Error() string {
	return "invalid password you fail because of me " + u.password
}
func userName() (string, error) {
	name := os.Args[1]
	data := usererror{
		name: name,
	}
	if data.name == "hello" {
		return data.name, fmt.Errorf("not able to validate user %w", data)
	}
	return name, nil
}

func userPass() (string, error) {
	password := os.Args[2]
	data := passerror{
		password: password,
	}
	if data.password == "mypass" {
		return data.password, fmt.Errorf("not able to validate password %w", data)
	}
	return password, nil
}

func main() {
	value, err := userName()
	da := usererror{}
	if err != nil {
		if errors.As(err, &da) {
			fmt.Println("be care ful next time")
		}
		fmt.Printf("not able to fetch user name ")
	}
	fmt.Println(value, err)
	pass, err := userPass()
	if err != nil {
		fmt.Printf("not able to fetch password ")
	}
	fmt.Println(pass, err)
}
