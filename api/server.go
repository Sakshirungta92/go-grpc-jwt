package api

import (
	"github.com/jinzhu/gorm"
	"authservice/pkg/auth"
)

type Server struct{
	DB		*gorm.DB
	Auth	auth.IAuth
}