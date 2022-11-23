package auth

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type AuthData struct {
	users map[string]string
}

func NewAuthData(filepath *string) (AuthData, error) {
	var users AuthData
	file, err := os.Open(*filepath)
	if err != nil {
		return users, fmt.Errorf("Failed open users file :%v", err)
	}
	defer file.Close()
	read, err := io.ReadAll(file)
	if err != nil {
		return users, fmt.Errorf("Failed reading file :%v", err)
	}
	err = yaml.Unmarshal(read, &users.users)
	if err != nil {
		return users, fmt.Errorf("Yaml unmarshalling error :%v", err)
	}
	return users, nil
}

func (usr *AuthData) CheckUserNameAndPassword(username, password string) bool {
	_, ok := usr.users[username]
	if !ok {
		return false
	}
	if usr.users[username] == password {
		return true
	}
	return false
}
