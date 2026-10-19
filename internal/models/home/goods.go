package home

type Goods struct {
	ID         string   `gorm:"primaryKey;size:32" json:"id"` // 字符串主键，删掉autoIncrement
	Name       string   `gorm:"type:varchar(100);not null" json:"name"`
	Picture    string   `gorm:"type:varchar(255)" json:"picture"`
	Desc       string   `gorm:"type:varchar(255)" json:"desc"`
	Price      string   `gorm:"type:varchar(255);" json:"price"`
	Discount   *float64 `gorm:"type:numeric(10,2)" json:"discount"`
	OrderNum   int64    `gorm:"default:0" json:"order_num"`
	CategoryID string   `gorm:"size:32;index;not null" json:"category_id"` // 字符串，关联category.id
}

func (Goods) TableName() string {
	return "goods"
}
