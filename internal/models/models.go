package models

import "time"

// Signup holds singup data
type Signup struct {
	Email             string
	EmailConfirmation string
	Password          string
}

// User is the user model
type User struct {
	ID                int
	Email             string
	EmailConfirmation string
	Password          string
	AccessLevel       int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// Product is the products model
type Product struct {
	ID          string
	Estilo      string
	Piel        string
	Suela       string
	Ca          bool
	Cantidad    int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Size is the sizes model
type Size struct {
	ID         string
	Size24     int
	Size24and5 int
	Size25     int
	Size25and5 int
	Size26     int
	Size26and5 int
	Size27     int
	Size27and5 int
	Size28     int
	Size28and5 int
	Size29     int
	Size29and5 int
	Size30     int
	Size30and5 int
	Size31     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Search is the search model that would be sent in a json object
type SearchBoot struct {
	Description string `json:"description"`
	Size24      int    `json:"size24"`
	Size24and5  int    `json:"size24And5"`
	Size25      int    `json:"size25"`
	Size25and5  int    `json:"size25And5"`
	Size26      int    `json:"size26"`
	Size26and5  int    `json:"size26And5"`
	Size27      int    `json:"size27"`
	Size27and5  int    `json:"size27And5"`
	Size28      int    `json:"size28"`
	Size28and5  int    `json:"size28And5"`
	Size29      int    `json:"size29"`
	Size29and5  int    `json:"size29And5"`
	Size30      int    `json:"size30"`
	Size30and5  int    `json:"size30And5"`
	Size31      int    `json:"size31"`
}

type ScanCode struct {
	Code string `json:"code"`
}

type ScanProduct struct {
	ID   string `json:"ID"`
	Size string `json:"Size"`
}
