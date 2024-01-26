package models

type DocumentCategory struct {
	DocumentId int `gorm:"type:int;not null" json:"document_id"`
	CategoryId int `gorm:"type:int;not null" json:"category_id"`
}

func (DocumentCategory) TableName() string {
	return "document_category"
}
