package db

type Product struct {
	Id             int64   `json:"id"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	ImgUrl         string  `json:"img_url"`
	AdditionalInfo string  `json:"additional_info"`
}

func (p Product) GetId() int64 {
	return p.Id
}
