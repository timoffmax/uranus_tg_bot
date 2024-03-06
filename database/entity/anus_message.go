package entity

import "gorm.io/gorm"

type AnusMessage struct {
	gorm.Model
	MessageId uint
}
