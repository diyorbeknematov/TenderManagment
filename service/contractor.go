package service

import (
	"fmt"
	"tender/model"
	"time"
)

func (s *Service) CreateBid(req *model.CreateBidInput) (*string, error) {
	if req.Price <= 0 {
		return nil, fmt.Errorf("invalid bid price: must be greater than zero")
	}
	deliveryTime, err := time.Parse("02-01-2006", req.DeliveryTime)
	if err != nil {
		return nil, fmt.Errorf("invalid delivery time format: %v", err)
	}
	currentTime := time.Now()
	if deliveryTime.Before(currentTime.Add(24 * time.Hour)) {
		return nil, fmt.Errorf("invalid delivery time: must be at least one day from today")
	}
	bidID, err := s.Storage.Contractor().CreateBid(*req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Bid creation failed for tender %s by contractor %s: %v", req.TenderID, req.ContractorID, err))
		return nil, fmt.Errorf("failed to create bid: %w", err)
	}
	return &bidID, nil
}

func (s *Service) GetTendersByFilters(req *model.GetTendersInput) (*[]model.Tender, error) {
	tenders, err := s.Storage.Contractor().GetTendersByFilters(*req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Fetching tenders failed: %v", err))
		return nil, err
	}
	return &tenders, nil
}

func (s *Service) GetBidsForTenderWithFilters(req *model.GetBidsInput) (*[]model.Bid, error) {
	bids, err := s.Storage.Contractor().GetBidsForTenderWithFilters(*req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Fetching bids failed: %v", err))
		return nil, err
	}
	return &bids, nil
}

func (s *Service) GetMyBidHistory(req *model.GetMyBidsInput) (*[]model.BidHistory, error) {
	history, err := s.Storage.Contractor().GetMyBidHistory(*req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Fetching bid history failed: %v", err))
		return nil, err
	}
	return &history, nil
}
