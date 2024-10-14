package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entry struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Dish        *string            `json:"dish,omitempty"`
	Fat         *float64           `json:"fat,omitempty"`
	Ingredients *string            `json:"ingredients,omitempty"`
	Calories    *string            `json:"calories,omitempty"`
}
