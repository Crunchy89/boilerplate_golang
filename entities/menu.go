package entities

type Menu struct {
	Title string `gorm:"size:100" json:"title"`
	Url   string `gorm:"size:255" json:"url"`
	Icon  string `gorm:"size:100" json:"icon"`
	Order int    `gorm:"default:1" json:"order"`
}
