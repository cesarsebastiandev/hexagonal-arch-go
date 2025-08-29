package domain

import "time"

type User struct {
	Name         string    `json:"name" binding:"required"`
	LastName     string    `json:"lastname" binding:"required"`
	Age          int       `json:"age" binding:"required"`
	CreationTime time.Time `json:"creation_time"`
}
