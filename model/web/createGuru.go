package web

type CreateGuru struct {
	Id_guru string `json:"id_guru"`
	Name    string `json:"name"`
	Status  *bool  `json:"status"`
}
