package services

type ThemeItem struct {
	ID             int         `json:"id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Name           string      `json:"name"`
	EntranceImg    string      `json:"entrance_img"`
	Extend         interface{} `json:"extend"`
	InternalTopImg interface{} `json:"internal_top_img"`
	TitleImg       interface{} `json:"title_img"`
	TplName        interface{} `json:"tpl_name"`
	Online         bool        `json:"online"`
}
type ThemeWithSpuItem struct {
	ThemeItem
	SpuList []SpuItem `json:"spu_list"`
}

type SpuItem struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle"`
	CategoryID     int    `json:"category_id"`
	RootCategoryID int    `json:"root_category_id"`
	Price          string `json:"price"`
	Img            string `json:"img"`
	ForThemeImg    string `json:"for_theme_img"`
	Description    string `json:"description"`
	DiscountPrice  string `json:"discount_price"`
	Tags           string `json:"tags"`
	IsTest         bool   `json:"is_test"`
	Online         bool   `json:"online"`
}
