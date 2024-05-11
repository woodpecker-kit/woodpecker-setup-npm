package npm_cfg

type VerdaccioLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type VerdaccioResponse struct {
	ErrorMsg string `json:"error"`
}

type VerdaccioLoginResponse struct {
	VerdaccioResponse
	Token    string `json:"token"`
	Username string `json:"username"`
}
