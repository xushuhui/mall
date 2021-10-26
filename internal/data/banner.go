package data

type Banner struct {
}
type Theme struct {
}

func GetBannerById(id int) (banner Banner, err error) {
	return
}
func GetBannerByName(name string) (banner Banner, err error) {
	return
}
func GetThemeWithSpu(name string) (theme Theme, err error) {
	return
}
func GetThemeByNames(names []string) (themes []Theme, err error) {
	return
}

type User struct {
}

func GetAccountUser(phone string) (user User, err error) {
	return
}

type Spu struct {
}

func GetSpuById(id int) (Spu Spu, err error) {
	return
}
func GetSpuByCategory(id int) (Spus []Spu, err error) {
	return
}
