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
	Pass                  string `json:"pass"`
	Mail                  string `json:"mail"`
	Theme                 string `json:"theme"`
	Signature             string `json:"signature"`
	Signature_format      string `json:"signature_format"`
	Created               int    `gorm:"not null;default:'0';" json:"Created"`
	Access                int    `gorm:"not null;default:'0';" json:"access"`
	Login                 int    `gorm:"not null;default:'0';" json:"login"`
	Status                int8   `gorm:"not nulldefault:'0';" json:"status"`
	Timezone              string `json:"timezone"`
	Language              string `json:"language"`
	Picture               int    `gorm:"not null;default:'0';" json:"picture"`
	Init                  string `json:"init"`
	Data                  string `json:"data"`
	Suspended             int    `gorm:"not null;default:'0'" json:"suspended"`
	Suspensions           int    `gorm:"not null;default:'0'" json:"suspension"`
	Subscribed            int    `gorm:"not null;default:'0'" json:"subscribed"`
	Point                 int    `gorm:"not null;default:'0'" json:"point"`
	Gamification_level    int    `gorm:"not null;default:'1'" json:"gamification_level"`
	Gamification_capacity int    `gorm:"not null;default:'0'" json:"gamification_capacity"`
	Webhook_transferred   int    `gorm:"not null;default:'0'" json:"webhook_transferred"`
	Owner                 int    `gorm:"not null;default:'0'" json:"owner"`
	Emoji_count           uint   `gorm:"not null;default:'0'" json:"emoji_count"`
	Emoji_details         string `gorm:"not null;default:''" json:"emoji_details"`
	Monetization          int    `gorm:"not null;default:'0'" json:"monetization"`
	Private               int    `gorm:"not null;default:'0'" json:"private"`
	Changed               int    `gorm:"not null;default:'0'" json:"changed"`
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
