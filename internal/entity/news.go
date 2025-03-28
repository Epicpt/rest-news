package entity

type News struct {
	ID         int64   `json:"Id"`
	Title      string  `json:"Title"`
	Content    string  `json:"Content"`
	Categories []int64 `json:"Categories"`
}

type NewsResponse struct {
	Success bool   `json:"Success"`
	News    []News `json:"News"`
}
