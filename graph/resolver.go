package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	//SessionStore map[string]model.Session
	//SeriesStore  map[string]model.Series
	//ShotStore    map[string]model.Shot
	DB *gorm.DB
}
