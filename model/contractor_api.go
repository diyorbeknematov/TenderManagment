package model

type CreateBid struct {
	Price        float64 `json:"price"`
	DeliveryTime string  `json:"delivery_time"`
	Comments     string  `json:"comments"`
}
