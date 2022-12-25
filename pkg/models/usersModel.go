package models

import (
	"github.com/jinzhu/gorm"
	"github.com/miladfallah/GoTarafdari/pkg/config"
)

var db *gorm.DB

type Users struct {
	gorm.Model
	Uid                   uint   `gorm:"primaryKey;not null;default:'0';" json:"uid"`
	Name                  string `gorm:"unique" json:"name"`
	pass                  string
	mail                  string
	theme                 string
	signature             string
	signature_format      string
	created               int  `gorm:"not null;default:'0';"`
	access                int  `gorm:"not null;default:'0';"`
	login                 int  `gorm:"not null;default:'0';"`
	status                int8 `gorm:"not nulldefault:'0';"`
	timezone              string
	language              string
	picture               int `gorm:"not null;default:'0';"`
	init                  string
	data                  string
	suspended             int    `gorm:"not null;default:'0'"`
	suspensions           int    `gorm:"not null;default:'0'"`
	subscribed            int    `gorm:"not null;default:'0'"`
	point                 int    `gorm:"not null;default:'0'"`
	gamification_level    int    `gorm:"not null;default:'1'"`
	gamification_capacity int    `gorm:"not null;default:'0'"`
	webhook_transferred   int    `gorm:"not null;default:'0'"`
	owner                 int    `gorm:"not null;default:'0'"`
	emoji_count           uint   `gorm:"not null;default:'0'"`
	emoji_details         string `gorm:"not null;default:''"`
	monetization          int    `gorm:"not null;default:'0'"`
	private               int    `gorm:"not null;default:'0'"`
	changed               int    `gorm:"not null;default:'0'"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Users{})
}

func (user *Users) CreateUser() *Users {
	db.NewRecord(user)
	db.Create(&user)
	return user
}
