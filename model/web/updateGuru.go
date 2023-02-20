package web

type UpdateGuru struct {
	Name   string `json:"name"`
	Status *bool  `json:"status"`
}
