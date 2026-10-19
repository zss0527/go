package home

type Category struct {
	ID       string  `gorm:"primaryKey;size:32" json:"id"` // 字符串主键，删掉autoIncrement
	Name     string  `gorm:"type:varchar(100);not null" json:"name"`
	Picture  string  `gorm:"type:varchar(255)" json:"picture"`
	ParentID *string `gorm:"size:32;index" json:"parent_id"` // 一级分类parent_id为nil
	Sort     int     `gorm:"default:0" json:"sort"`

	// GORM关联标记
	Children []Category `gorm:"foreignKey:ParentID;references:ID" json:"children"`
	Goods    []Goods    `gorm:"foreignKey:CategoryID;references:ID" json:"goods"`
}

func (Category) TableName() string {
	return "category"
}
