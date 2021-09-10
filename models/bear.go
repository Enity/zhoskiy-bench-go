package models

type Bear struct {
	Id              uint   `gorm:"primarykey" json:"id"`
	Name            string `json:"name" validate:"required"`
	KdRatio         string `json:"kdRatio" validate:"required"`
	LoveToSuckCocks bool   `json:"loveToSuckCocks" validate:"required"`
}
