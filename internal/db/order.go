package db

type Order struct {
	Id        int64 `json:"id"`
	CourierId int64 `json:"courier_id"`
}

func (o Order) GetId() int64 {
	return o.Id
}
