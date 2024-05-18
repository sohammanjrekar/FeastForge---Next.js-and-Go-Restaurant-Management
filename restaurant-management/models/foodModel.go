package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required",min=2,max=100`
	Price      *float64           `json:"price" validate:"required"`
	Food_image *float64           `json:"food_image" validate:"required"`

	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Food_id    string    `json:"food_id"`
	Menu_id    *float64  `json:"menu_id" validate:"required"`
}
