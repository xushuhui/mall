package request

type Login struct {
	Phone    string `json:"phone" validate:"required,mobile" comment:"手机号"`
	Password string `json:"password" validate:"required" comment:"密码"`
}
