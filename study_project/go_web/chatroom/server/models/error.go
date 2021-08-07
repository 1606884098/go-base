package models

import "errors"

var (
	ERROR_USER_NOEXISTS = errors.New("用户不存在...")
	ERROR_USER_EXISTS   = errors.New("用户已经存在...")
	ERROR_USER_PWD      = errors.New("密码不正确...")
)
