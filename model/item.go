package model

type Items struct {
	ID          string `json:"id" gorm:"size:32;unique"`
	Description string `json:"description" gorm:"size:64;unique"`
	Point       int    `json:"point" validate:"omitempty"`
	Report      int    `json:"report" validate:"omitempty"`
}