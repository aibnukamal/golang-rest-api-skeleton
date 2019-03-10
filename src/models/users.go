package models

import "github.com/jinzhu/gorm"

// Users data model
type Users struct {
	gorm.Model
	Firstname        string `json:"firstname" gorm:"column:firstname;size:255"`
	Lastname         string `json:"lastname" gorm:"column:lastname;size:255"`
	Username         string `json:"username" gorm:"column:username;size:15"`
	Email            string `json:"email" gorm:"column:email;size:255"`
	Password         string `json:"password" gorm:"column:password;size:255"`
	Photo            string `json:"photo" gorm:"column:photo;size:255"`
	Address          string `json:"address" gorm:"column:address;size:255"`
	IDCardNumber     string `json:"id_card_number" gorm:"column:id_card_number;size:100"`
	IDCardImage      string `json:"id_card_image" gorm:"column:id_card_image;size:255"`
	IDCardImagePhoto string `json:"id_card_image_photo" gorm:"column:id_card_image_photo;size:255"`
	Status           uint   `json:"status" gorm:"column:status"`
}
