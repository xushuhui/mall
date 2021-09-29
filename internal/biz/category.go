package biz

type CategoryItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsRoot   bool   `json:"is_root"`
	Img      string `json:"img"`
	ParentID int    `json:"parent_id"`
	Index    int    `json:"index"`
}

type GridItem struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Img            string `json:"img"`
	Name           string `json:"name"`
	CategoryID     int    `json:"category_id"`
	RootCategoryID int    `json:"root_category_id"`
}
