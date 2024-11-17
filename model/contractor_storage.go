package model

import "time"

type CreateBidInput struct {
	TenderID     string  `json:"tender_id"`
	ContractorID string  `json:"contractor_id"`
	Price        float64 `json:"price"`
	DeliveryTime string  `json:"delivery_time"`
	Comments     string  `json:"comments"`
}

type GetTendersInput struct {
	Status string `json:"status"`
}

type GetBidsInput struct {
	TenderID        string  `json:"tender_id"`
	MaxPrice        float64 `json:"max_price"`
	MaxDeliveryTime string  `json:"max_delivery_time"`
}

type Bid struct {
	ID           string    `json:"id"`
	TenderID     string    `json:"tender_id"`
	ContractorID string    `json:"contractor_id"`
	Price        float64   `json:"price"`
	DeliveryTime time.Time `json:"delivery_time"`
	Comments     string    `json:"comments"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type GetMyBidsInput struct {
	UserID string `json:"user_id"`
}

type BidHistory struct {
	Bid
	TenderTitle    string    `json:"tender_title"`
	TenderDeadline time.Time `json:"tender_deadline"`
}
