package domain

type DataList struct {
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	LinkDetails string `json:"link_details"`
}

type CommandRepository interface {
	Find(keyword string) (*[]DataList, error)
}

type CommandService interface {
	Ping() (string, error)
	Help() (string, error)
	Find(keyword string) (string, error)
}
