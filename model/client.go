package model

type Error struct {
	Message string `json:"message"`
}

type CreateTenderReq struct {
	ClientId    string  `json:"client_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Diadline    string  `json:"diadline"`
	Budget      float64 `json:"budget"`
}

type CreateTenderResp struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type GetAllTendersReq struct {
	ClientId string `json:"client_id"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}

type Tender struct {
	Id          string  `json:"id"`
	ClientId    string  `json:"client_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Diadline    string  `json:"diadline"`
	Budget      float64 `json:"budget"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
}

type GetAllTendersResp struct {
	Tenders []Tender `json:"tenders"`
	Count   int      `json:"count"`
}

type UpdateTenderReq struct {
	Id          string  `json:"id"`
	ClientId    string  `json:"client_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Diadline    string  `json:"diadline"`
	Budget      float64 `json:"budget"`
	Status      string  `json:"status"`
}

type UpdateTenderResp struct {
	Message string `json:"message"`
}

type DeleteTenderReq struct {
	Id       string `json:"id"`
	ClientId string `json:"client_id"`
}

type DeleteTenderResp struct {
	Message string `json:"message"`
}

type GetTenderBidsReq struct {
	ClientId   string  `json:"client_id" form:"client_id"`
	TenderId   string  `json:"tender_id"`
	StartPrice float64 `json:"start_price" form:"start_price"`
	EndPrice   float64 `json:"end_price" form:"end_price"`
	StartDate  string  `json:"start_date" form:"start_date"`
	EndDate    string  `json:"end_date" form:"end_date"`
	Limit      int     `json:"limit"`
	Page       int     `json:"page"`
}

type GetTenderBidsResp struct {
	Bids  []Bid `json:"bids"`
	Count int   `json:"count"`
}

type BidAwardedReq struct {
	ClientId string `json:"client_id"`
	TenderId string `json:"tender_id"`
	BidId    string `json:"bid_id"`
	Status   string `json:"status"`
}

type BidAwardedResp struct {
	Status bool `json:"status"`
}
