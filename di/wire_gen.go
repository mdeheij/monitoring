// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"waarnemer/repository"
)

// Injectors from wire.go:

func InitializeCheckRepository() *repository.CheckRepository {
	checkRepository := repository.NewCheckRepository()
	return checkRepository
}