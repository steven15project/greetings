package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, errors) {
	if name == "" {
		return "", errors.New("nama kosong")
	}
	message := fmt.Sprintf("Hi, %v	. Welcome!", name)
	return message, nil
}
