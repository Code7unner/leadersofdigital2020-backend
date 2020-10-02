package db

type Store struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}

func (s Store) GetId() int64 {
	return s.Id
}