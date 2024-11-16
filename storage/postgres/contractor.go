package postgres

import (
	"database/sql"
	"fmt"
	"tender/model"
)

func CreateBid(db *sql.DB, input model.CreateBidInput) (int, error) {
	var bidID int
	query := `
        INSERT INTO bids (tender_id, contractor_id, price, delivery_time, comments, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
        RETURNING id;
    `
	err := db.QueryRow(query, input.TenderID, input.ContractorID, input.Price, input.DeliveryTime, input.Comments, input.Status).Scan(&bidID)
	if err != nil {
		return 0, fmt.Errorf("failed to create bid: %w", err)
	}
	return bidID, nil
}

func GetTendersByStatus(db *sql.DB, input model.GetTendersInput) ([]model.Tender, error) {
	query := `
        SELECT id, client_id, title, description, deadline, budget, status, created_at, updated_at, deleted_at
        FROM tenders
        WHERE status = $1
          AND (deleted_at IS NULL)
        ORDER BY deadline ASC;
    `
	rows, err := db.Query(query, input.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tenders: %w", err)
	}
	defer rows.Close()

	var tenders []model.Tender
	for rows.Next() {
		var tender model.Tender
		if err := rows.Scan(&tender.ID, &tender.ClientID, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.CreatedAt, &tender.UpdatedAt, &tender.DeletedAt); err != nil {
			return nil, fmt.Errorf("failed to scan tender: %w", err)
		}
		tenders = append(tenders, tender)
	}
	return tenders, nil
}

func GetBidsForTender(db *sql.DB, input model.GetBidsInput) ([]model.Bid, error) {
	query := `
        SELECT id, tender_id, contractor_id, price, delivery_time, comments, status, created_at, updated_at, deleted_at
        FROM bids
        WHERE tender_id = $1
          AND price <= $2
          AND delivery_time <= $3
          AND (deleted_at IS NULL)
        ORDER BY price ASC, delivery_time ASC;
    `
	rows, err := db.Query(query, input.TenderID, input.MaxPrice, input.MaxDeliveryTime)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bids: %w", err)
	}
	defer rows.Close()

	var bids []model.Bid
	for rows.Next() {
		var bid model.Bid
		if err := rows.Scan(&bid.ID, &bid.TenderID, &bid.ContractorID, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status, &bid.CreatedAt, &bid.UpdatedAt, &bid.DeletedAt); err != nil {
			return nil, fmt.Errorf("failed to scan bid: %w", err)
		}
		bids = append(bids, bid)
	}
	return bids, nil
}

func GetMyBidHistory(db *sql.DB, input model.GetMyBidsInput) ([]model.BidHistory, error) {
	query := `
        SELECT b.id, b.tender_id, b.contractor_id, b.price, b.delivery_time, b.comments, b.status, 
               b.created_at, b.updated_at, b.deleted_at, t.title AS tender_title, t.deadline AS tender_deadline
        FROM bids b
        JOIN tenders t ON b.tender_id = t.id
        WHERE b.contractor_id = $1
          AND (b.deleted_at IS NULL)
        ORDER BY b.created_at DESC;
    `
	rows, err := db.Query(query, input.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bid history: %w", err)
	}
	defer rows.Close()

	var history []model.BidHistory
	for rows.Next() {
		var bid model.BidHistory
		if err := rows.Scan(&bid.ID, &bid.TenderID, &bid.ContractorID, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status, &bid.CreatedAt, &bid.UpdatedAt, &bid.DeletedAt, &bid.TenderTitle, &bid.TenderDeadline); err != nil {
			return nil, fmt.Errorf("failed to scan bid history: %w", err)
		}
		history = append(history, bid)
	}
	return history, nil
}
