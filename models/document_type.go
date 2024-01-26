package models

type DocumentType struct {
	TypeId   int    `gorm:"primary_key;AUTO_INCREMENT" json:"type_id"`
	TypeName string `gorm:"type:varchar(255);not null" json:"type_name"`
}

func (DocumentType) TableName() string {
	return "document_type"
}
