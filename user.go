package teamcity10

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type Triggered struct {
	Type string `json:"type"`
	User User   `json:"user"`
}
