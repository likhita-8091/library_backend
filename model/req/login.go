package req

type Login struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type ChangePasswordReq struct {
	Password    string `json:"password"`     // 密码
	NewPassword string `json:"new_password"` // 新密码
}
