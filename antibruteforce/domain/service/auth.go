package service

import "github.com/teploff/otus/antibruteforce/domain/entity"

type AuthService interface {
	LogIn(credentials entity.Credentials, ip string) (bool, error)
}
