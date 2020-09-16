package teamcity10

type Build struct {
	ID        int       `json:"id"`
	Type      string    `json:"buildTypeId"`
	Number    string    `json:"number"`
	Status    string    `json:"status"`
	State     string    `json:"state"`
	Branch    string    `json:"branchName"`
	EndedAt   string    `json:"finishDate"`
	Triggered Triggered `json:"triggered"`
}
