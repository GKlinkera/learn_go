package function

type Infouser struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Tianqi struct {
	Rain  string
	Cloud string
}
