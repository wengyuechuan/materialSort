package models

type Document struct {
	DocumentId  int    `gorm:"primary_key;AUTO_INCREMENT" json:"document_id"`
	TypeId      int    `gorm:"type:int;not null" json:"type_id"`
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	FilePath    string `gorm:"type:varchar(255)" json:"file_path"`
	FileMd5     string `gorm:"type:varchar(32)" json:"file_md5"`
	UploadDate  string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"upload_date"`
	UpdateDate  string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_date"`
	UserId      int    `gorm:"type:int" json:"user_id"`
}

func (Document) TableName() string {
	return "document"
}
