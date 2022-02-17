package reststructs

type Login struct {
	APIKey string `json:"api_key"`
	Pin    string `json:"pin"`
}

type LoginToken struct {
	Token string
}
