package entity

type Info struct {
	ID          int32  `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	Title       string `gorm:"column:title; type:varchar(255); not null"`
	Key         string `gorm:"column:key; type:varchar(255); not null"`
	Categories  string `gorm:"column:categories; type:varchar(100); not null"`
	Type        string `gorm:"column:type; type:varchar(50); not null"`
	Language    string `gorm:"column:language; type:varchar(10); default:Null"`
	Description string `gorm:"column:description;default:Null"`
	Image       string `gorm:"column:image; type:varchar(255); default:Null"`
	Version     string `gorm:"column:version; type:varchar(100); not null"`
	CreatedAt   int64  `gorm:"column:created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at"`
	CreatedBy   string `gorm:"column:created_by"`
	UpdatedBy   string `gorm:"column:updated_by"`
}

func (Info) TableName() string {
	return "infos"
}
