package entities

import "time"

// deklarasi key table database
type Category struct {
	Id        uint
	Name      string
	CreateAt time.Time
	UpdateAt time.Time
}