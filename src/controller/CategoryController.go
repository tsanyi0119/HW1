package controller

type Category struct {
    ID   uint   `gorm:"primaryKey" json:"id"`
    Name string `json:"name"`
}