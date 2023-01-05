package filter

import . "go-service/internal/model"

type MovieFilter struct {
	Id      string `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id" validate:"required,max=40"`
	Name    string `json:"name" gorm:"column:name" bson:"name" dynamodbav:"name" firestore:"name" validate:"required,name,max=100"`
	Watched bool   `json:"watched" validate:"required"`
}

type ResultMovie struct {
	List  []Movie `mapstructure:"list" json:"list,omitempty" gorm:"column:list" bson:"list,omitempty" dynamodbav:"list,omitempty" firestore:"list,omitempty"`
	Total int64   `mapstructure:"total" json:"total,omitempty" gorm:"column:total" bson:"total,omitempty" dynamodbav:"total,omitempty" firestore:"total,omitempty"`
}
