package repository

import "github.com/DeLuci/inventory-system/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertUser(res models.User) error
	Authenticate(email, password string) (int, string, error)
	InsertFinishedProduct(res models.Product) error
	InsertNewProduct(res models.ScanProduct) error
	SearchProduct(searchName string) ([]models.SearchBoot, error)
}
