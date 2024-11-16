package model

import "time"

type CreateBidInput struct {
	TenderID     string
	ContractorID string
	Price        float64
	DeliveryTime string
	Comments     string
}

type GetTendersInput struct {
	Status string `json:"status"` // Filter by status
}

type GetBidsInput struct {
	TenderID        string
	MaxPrice        float64
	MaxDeliveryTime string
}

type Bid struct {
	ID           string
	TenderID     string
	ContractorID string
	Price        float64
	DeliveryTime time.Time
	Comments     string
	Status       string
	CreatedAt    time.Time
}

type GetMyBidsInput struct {
	UserID string
}

type BidHistory struct {
	Bid
	TenderTitle    string
	TenderDeadline time.Time
}
