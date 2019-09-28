package model

import "errors"
var (

	ErrUserNotExist = errors.New("user not exist")
	ErrInvalidPasswd = errors.New("passwd is invalid")

	ErrInvalidParams = errors.New("params is invalid")

	ErrUserExist = errors.New("user is exist")


)
