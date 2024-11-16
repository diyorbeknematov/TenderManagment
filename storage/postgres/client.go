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
					client_id, title, description, deadline, budget, status)
				VALUES
					($1, $2, $3, $4, $5, $6)`

	_, err := C.DB.Exec(query, id, req.Diadline, req.Description, req.Diadline, req.Diadline, req.Budget, "open")
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
					id, title, description, diadline, budget, status
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
		err := rows.Scan(&tender.Id, &tender.Title, &tender.Description, &tender.Diadline, &tender.Budget, &tender.Status)
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
					title = $1, description = $2, diadline = $3, budget = $4, status = $5, updated_at = $8
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
				UPDATE Tensers SET 
					deleted_at = $1
				WHERE 
					id = $2 AND client_id = $2 AND deleted_at IS NULL`
	_, err := C.DB.Exec(query, time.Now(), req.Id, req.ClientId)
	if err != nil {
		C.Log.Error(fmt.Sprintf("Tender o'chirilmadi: %v", err))
		return nil, err
	}

	return &model.DeleteTenderResp{
		Message: "Tender muvaffaqiyatli o'chirildi",
	}, nil
}

func (C *clientImpl) GetTenderBids(req *)
