package database

type User struct {
	Base

	Name    string `gorm:"uniqueIndex"`
	PWDHash string
	IsAdmin bool
}
