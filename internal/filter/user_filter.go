package filter

import . "go-service/internal/model"

type UserFilter struct {
	Id        string `mapstructure:"id" json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id" match:"equal" validate:"required,max=40"`
	Username  string `mapstructure:"username" json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" match:"prefix" validate:"required,username,max=100"`
	Email     string `mapstructure:"email" json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" match:"prefix" validate:"email,max=100"`
	Phone     string `mapstructure:"phone" json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" validate:"required,phone,max=18"`
	PageIndex int64  `mapstructure:"pageIndex" json:"pageIndex,omitempty" gorm:"column:pageIndex" bson:"pageIndex,omitempty" dynamodbav:"pageIndex,omitempty" firestore:"pageIndex,omitempty"`
	PageSize  int64  `mapstructure:"pageSize" json:"pageSize,omitempty" gorm:"column:pageSize" bson:"pageSize,omitempty" dynamodbav:"pageSize,omitempty" firestore:"pageSize,omitempty"`
}

type Result struct {
	List  []User `mapstructure:"list" json:"list,omitempty" gorm:"column:list" bson:"list,omitempty" dynamodbav:"list,omitempty" firestore:"list,omitempty"`
	Total int64  `mapstructure:"total" json:"total,omitempty" gorm:"column:total" bson:"total,omitempty" dynamodbav:"total,omitempty" firestore:"total,omitempty"`
}
