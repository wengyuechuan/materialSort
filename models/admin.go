package models

type Admin struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"type:varchar(255);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

func (Admin) TableName() string {
	return "admin"
}
