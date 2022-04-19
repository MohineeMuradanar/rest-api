package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model
type Student struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name,omitempty"`
	City            string             `json:"city,omitempty" bson:"city,omitempty"`
	Country         string             `json:"country,omitempty" bson:"country,omitempty"`
	Course          string             `json:"course,omitempty" bson:"course,omitempty"`
	YearOfAdmission int                `json:"year_of_admission,omitempty" bson:"yearofadmission,omitempty"`
}
