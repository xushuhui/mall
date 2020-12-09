package model

type Category struct {
	BaseModel
	Name        string `gorm:"column:name;" json:"name"`
	Description string `gorm:"column:description;" json:"description"`
	IsRoot      int    `gorm:"column:is_root;default:'0'" json:"is_root"`
	ParentId    int    `gorm:"column:parent_id;" json:"parent_id"`
	Img         string `gorm:"column:img;" json:"img"`
	Index       int    `gorm:"column:index;" json:"index"`
	Online      int    `gorm:"column:online;default:'1'" json:"online"`
	Level       int    `gorm:"column:level;" json:"level"`
}
type GridCategory struct {
	BaseModel

	Title string `json:"title"`
	Img   string `json:"img"`
	Name  string `json:"name"`

	CategoryID     int32 `json:"category_id"`
	RootCategoryID int32 `json:"root_category_id"`
}
