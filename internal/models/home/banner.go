package home

type Banner struct {
	ID      string `gorm:"primaryKey;size:32" json:"id"`
	ImgURL  string `gorm:"type:varchar(255);" json:"img_url"`
	HrefURL string `gorm:"type:varchar(255);" json:"href_url"`
	Type    string `gorm:"type:varchar(10);" json:"type"`
}

func (Banner) TableName() string {
	return "banner"
}
