package npm_cfg

type VerdaccioLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type VerdaccioLoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}
