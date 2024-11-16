package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"tender/model"
)

func CreateBid(db *sql.DB, input model.CreateBidInput) (int, error) {
	var bidID int
	query := `
        INSERT INTO bids (tender_id, contractor_id, price, delivery_time, comments, status)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `
	err := db.QueryRow(query, input.TenderID, input.ContractorID, input.Price, input.DeliveryTime, input.Comments, input.Status).Scan(&bidID)
	if err != nil {
		return 0, fmt.Errorf("failed to create bid: %w", err)
	}
	return bidID, nil
}

func GetTendersByFilters(db *sql.DB, input model.GetTendersInput) ([]model.Tender, error) {
	query := `
        SELECT id, client_id, title, description, deadline, budget, status, created_at
        FROM tenders
        WHERE deleted_at IS NULL`
	var filters []string
	var params []interface{}
	paramIndex := 1

	if input.Status != "" {
		filters = append(filters, fmt.Sprintf("status=$%d", paramIndex))
		params = append(params, input.Status)
		paramIndex++
	}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}

	query += " ORDER BY deadline ASC;"

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tenders: %w", err)
	}
	defer rows.Close()

	var tenders []model.Tender
	for rows.Next() {
		var tender model.Tender
		if err := rows.Scan(&tender.ID, &tender.ClientID, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan tender: %w", err)
		}
		tenders = append(tenders, tender)
	}

	return tenders, nil
}

func GetBidsForTenderWithFilters(db *sql.DB, input model.GetBidsInput) ([]model.Bid, error) {
	query := `
        SELECT id, tender_id, contractor_id, price, delivery_time, comments, status, created_at
        FROM bids
        WHERE deleted_at IS NULL AND tender_id = $1`
	var filters []string
	var params []interface{}
	params = append(params, input.TenderID)
	paramIndex := 2

	if input.MaxPrice > 0 {
		filters = append(filters, fmt.Sprintf("price <= $%d", paramIndex))
		params = append(params, input.MaxPrice)
		paramIndex++
	}
	if input.MaxDeliveryTime != "" {
		filters = append(filters, fmt.Sprintf("delivery_time <= $%d", paramIndex))
		params = append(params, input.MaxDeliveryTime)
		paramIndex++
	}

	if len(filters) > 0 {
		query += " AND " + fmt.Sprintf("%s", filters[0])
		for _, filter := range filters[1:] {
			query += " AND " + filter
		}
	}

	query += " ORDER BY price ASC, delivery_time ASC;"

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bids: %w", err)
	}
	defer rows.Close()

	var bids []model.Bid
	for rows.Next() {
		var bid model.Bid
		if err := rows.Scan(&bid.ID, &bid.TenderID, &bid.ContractorID, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status, &bid.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan bid: %w", err)
		}
		bids = append(bids, bid)
	}
	return bids, nil
}

func GetMyBidHistory(db *sql.DB, input model.GetMyBidsInput) ([]model.BidHistory, error) {
	query := `
        SELECT b.id, b.tender_id, b.contractor_id, b.price, b.delivery_time, b.comments, b.status, 
               b.created_at, t.title AS tender_title, t.deadline AS tender_deadline
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
		if err := rows.Scan(&bid.ID, &bid.TenderID, &bid.ContractorID, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status, &bid.CreatedAt, &bid.TenderTitle, &bid.TenderDeadline); err != nil {
			return nil, fmt.Errorf("failed to scan bid history: %w", err)
		}
		history = append(history, bid)
	}
	return history, nil
}
