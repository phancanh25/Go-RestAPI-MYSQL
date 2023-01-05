package model

type Movie struct {
	Id      string `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id" validate:"required,max=40"`
	Name    string `json:"name" gorm:"column:name" bson:"name" dynamodbav:"name" firestore:"name" validate:"required,name,max=100"`
	Watched bool   `json:"watched" validate:"required"`
}
