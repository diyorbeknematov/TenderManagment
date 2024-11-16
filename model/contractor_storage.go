package model

import "time"

type CreateBidInput struct {
	TenderID     int
	ContractorID int
	Price        float64
	DeliveryTime string
	Comments     string
}

type GetTendersInput struct {
	Status string `json:"status"` // Filter by status
}

type GetBidsInput struct {
	TenderID        int
	MaxPrice        float64
	MaxDeliveryTime string
}

type Bid struct {
	ID           int
	TenderID     int
	ContractorID int
	Price        float64
	DeliveryTime time.Time
	Comments     string
	Status       string
	CreatedAt    time.Time
}

type GetMyBidsInput struct {
	UserID int
}

type BidHistory struct {
	Bid
	TenderTitle    string
	TenderDeadline time.Time
}
