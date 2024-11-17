package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"tender/model"
	"time"

	"github.com/google/uuid"
)

type ClientRepo interface {
	CreateTender(req *model.CreateTenderReq) (*model.CreateTenderResp, error)
	GetAllTenders(req *model.GetAllTendersReq) (*model.GetAllTendersResp, error)
	UpdateTender(req *model.UpdateTenderReq) (*model.UpdateTenderResp, error)
	DeleteTender(req *model.DeleteTenderReq) (*model.DeleteTenderResp, error)
	GetTenderBids(req *model.GetTenderBidsReq) (*model.GetTenderBidsResp, error)
	SubmitBit(req *model.SubmitBitReq) (*model.SubmitBitResp, error)
	AwardTender(req *model.AwardTenderReq) (*model.AwardTenderResp, error)
	GetUserByTebderId(tenderId string) (string, error)
	CloseExpiredTenders() error
}

type clientImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewClientRepo(db *sql.DB, logger *slog.Logger) ClientRepo {
	return &clientImpl{
		DB:  db,
		Log: logger,
	}
}

func (C *clientImpl) CreateTender(req *model.CreateTenderReq) (*model.CreateTenderResp, error) {
	id := uuid.NewString()

	query := `
				INSERT INTO Tenders(
					id, client_id, title, description, deadline, budget, status)
				VALUES
					($1, $2, $3, $4, $5, $6, $7)`

	_, err := C.DB.Exec(query, id, req.ClientId, req.Title, req.Description, req.Diadline, req.Budget, "open")
	if err != nil {
		C.Log.Error(fmt.Sprintf("Ma'lumotlarni databazaga saqlashda xatolik: %v", err))
		return nil, err
	}

	return &model.CreateTenderResp{
		Id:        id,
		Message:   "Tender muvaffaqiyatli yaratildi",
		CreatedAt: time.Now().String(),
	}, nil
}

func (C *clientImpl) GetAllTenders(req *model.GetAllTendersReq) (*model.GetAllTendersResp, error) {
	var tenders = []model.Tender{}
	offset := (req.Page - 1) * req.Limit
	var count int

	query := `
				SELECT 
					id, client_id, title, description, deadline, budget, status, created_at
				FROM 
					Tenders
				WHERE 
					client_id = $1 AND deleted_at IS NULL 
				LIMIT $2
				OFFSET $3`

	rows, err := C.DB.Query(query, req.ClientId, req.Limit, offset)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tender model.Tender
		err := rows.Scan(&tender.Id, &tender.ClientId, &tender.Title, &tender.Description, &tender.Diadline, &tender.Budget, &tender.Status, &tender.CreatedAt)
		if err != nil {
			C.Log.Error(fmt.Sprintf("Ma'lumotlarni o'zlashtirishda xatolik: %v", err))
			return nil, err
		}
		tenders = append(tenders, tender)
	}

	query = `
				SELECT
					COUNT(*)
				FROM 
					Tenders
				WHERE
					client_id = $1 AND deleted_at IS NULL`
	err = C.DB.QueryRow(query, req.ClientId).Scan(&count)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Tenderlar sonini olishda xatolik: %v", err))
		return nil, err
	}

	return &model.GetAllTendersResp{
		Tenders: tenders,
		Count:   count,
	}, nil
}

func (C *clientImpl) UpdateTender(req *model.UpdateTenderReq) (*model.UpdateTenderResp, error) {
	query := `
				UPDATE Tenders SET
					title = $1, description = $2, deadline = $3, budget = $4, status = $5, updated_at = $8
				WHERE 
					client_id = $6 AND id = $7 AND deleted_at IS NULL`
	_, err := C.DB.Exec(query, req.Title, req.Description, req.Diadline, req.Budget, req.Status, req.ClientId, req.Id, time.Now())
	if err != nil {
		C.Log.Error(fmt.Sprintf("Ma'lumotlarni yangilashda xatolik: %v", err))
		return nil, err
	}

	return &model.UpdateTenderResp{
		Message: "Tender muvaffaqiyatli yangilandi",
	}, nil
}

func (C *clientImpl) DeleteTender(req *model.DeleteTenderReq) (*model.DeleteTenderResp, error) {
	query := `
				UPDATE Tenders SET 
					deleted_at = $1
				WHERE 
					id = $2 AND client_id = $3 AND deleted_at IS NULL`
	_, err := C.DB.Exec(query, time.Now(), req.Id, req.ClientId)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Tender o'chirilmadi: %v", err))
		return nil, err
	}

	return &model.DeleteTenderResp{
		Message: "Tender muvaffaqiyatli o'chirildi",
	}, nil
}

func (C *clientImpl) GetTenderBids(req *model.GetTenderBidsReq) (*model.GetTenderBidsResp, error) {
	var bids []model.Bid
	var param []interface{}
	var count int
	query := `
				SELECT 
					id, tender_id, contractor_id, price, delivery_time, comments, status, created_at
				FROM 
					bids
				WHERE
					tender_id = $1 AND deleted_at IS NULL`
	count_query := `
						SELECT 
							COUNT(*)
						FROM 
							bids
						WHERE 
							tender_id = $1 AND deleted_at IS NULL`
	param = append(param, req.TenderId)
	if len(req.StartDate) > 0 {
		param = append(param, req.StartDate)
		query += fmt.Sprintf(" AND delivery_time >= $%v", len(param))
		count_query += fmt.Sprintf(" AND delivery_time >= $%v", len(param))
	}
	if len(req.EndDate) > 0 {
		param = append(param, req.EndDate)
		query += fmt.Sprintf(" AND delivery_time <= $%v", len(param))
		count_query += fmt.Sprintf(" AND delivery_time <= $%v", len(param))
	}
	if req.StartPrice >= 0.0 {
		param = append(param, req.StartPrice)
		query += fmt.Sprintf(" AND price >= $%v", len(param))
		count_query += fmt.Sprintf(" AND price >= $%v", len(param))
	}
	if req.EndPrice > 0.0 {
		param = append(param, req.EndPrice)
		query += fmt.Sprintf(" AND price <= $%v", len(param))
		count_query += fmt.Sprintf(" AND price <= $%v", len(param))
	}
	query += fmt.Sprintf(" LIMIT %v", req.Limit)
	offset := (req.Page - 1) * req.Limit
	query += fmt.Sprintf(" OFFSET %v", offset)

	rows, err := C.DB.Query(query, param...)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bid model.Bid
		err = rows.Scan(&bid.ID, &bid.TenderID, &bid.ContractorID, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status, &bid.CreatedAt)
		if err != nil {
			C.Log.Error(fmt.Sprintf("Bid ma'lumotlarini o'zlashtirishda xatolik: %v", err))
			return nil, err
		}
		bids = append(bids, bid)
	}

	err = C.DB.QueryRow(count_query, param...).Scan(&count)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Bidlar sonini olishda xatolik: %v", err))
		return nil, err
	}

	return &model.GetTenderBidsResp{
		Bids:  bids,
		Count: count,
	}, nil
}

func (C *clientImpl) SubmitBit(req *model.SubmitBitReq) (*model.SubmitBitResp, error) {
	query := `
				UPDATE bids SET
					status = $1 
				WHERE 
					id = $2 AND deleted_at IS NULL`
	_, err := C.DB.Exec(query, req.Status, req.BidId)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Bid statusi yangilanmadi: %v", err))
		return nil, err
	}

	return &model.SubmitBitResp{
		Status: true,
	}, nil
}

func (C *clientImpl) AwardTender(req *model.AwardTenderReq) (*model.AwardTenderResp, error) {
	tr, err := C.DB.Begin()
	if err != nil {
		C.Log.Error(fmt.Sprintf("Tranzaksiya yaratishda xatolik: %v", err))
		return nil, err
	}
	defer tr.Commit()

	query := `
				UPDATE bids SET 
					status = $1
				WHERE
					tender_id = $2 AND id = $3 AND deleted_at IS NULL`
	_, err = tr.Exec(query, "award", req.TenderId, req.BidId)
	if err != nil{
		C.Log.Error(fmt.Sprintf("Contractorga award berishda xatolik: %v", err))
		tr.Rollback()
		return nil, err
	}

	query = `
				UPDATE bids SET
					status = $1
				WHERE
					tender_id = $2 AND id != $3 AND deleted_at IS NULL`
	_, err = tr.Exec(query, "fail", req.TenderId, req.BidId)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Contractorlargaga fail berishda xatolik: %v", err))
		tr.Rollback()
		return nil, err
	}

	query = `
				UPDATE tenders SET 
					status = $1
				WHERE 
					client_id = $2 AND id = $3 AND deleted_at IS NULL`
	_, err = tr.Exec(query, "awarded", req.ClientId, req.TenderId)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Tender statusini o'zgartirishda xatolik: %v", err))
		tr.Rollback()
		return nil, err
	}

	return &model.AwardTenderResp{
		Status: true,
	}, nil
}

func (C *clientImpl) GetUserByTebderId(tenderId string) (string, error) {
	var userId string
	query := `
				SELECT 
					client_id
				FROM 
					tenders
				WHERE
					id = $1 AND deleted_at IS NULL`
	err := C.DB.QueryRow(query, tenderId).Scan(&userId)
	if err != nil {
		C.Log.Error(fmt.Sprintf("TenderId bo'yicha ClientIdisini olishda xatolik: %v", err))
		return "", err
	}

	return userId, nil
}

func (s *clientImpl) CloseExpiredTenders() error {
	query := `
		UPDATE tenders
		SET status = 'close', updated_at = NOW()
		WHERE status = 'open' AND deadline < NOW();
	`
	result, err := s.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to update expired tenders: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Expired tenders updated: %d\n", rowsAffected)
	return nil
}
