package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quiz struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name"`
	Questions []QuizQuestion     `json:"questions"`
}

type QuizQuestion struct {
	Id      string       `json:"id" bson:"id"`
	Name    string       `json:"name" bson:"name"`
	Time    int          `json:"time" bson:"time"`
	Image   string       `json:"image" bson:"image,omitempty"`
	Choices []QuizChoice `json:"choices" bson:"choices"`
}

type QuizChoice struct {
	Id      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Correct bool   `json:"correct" bson:"correct"`
}
