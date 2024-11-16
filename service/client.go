package service

import (
	"fmt"
	"log/slog"
	"tender/model"
	"tender/storage"
)

type Service struct {
	Storage storage.Storage
	Log     *slog.Logger
}

func NewService(storage storage.Storage, logger *slog.Logger)*Service{
	return &Service{
		Storage: storage,
		Log: logger,
	}
}

func(S *Service) CreateTender(req *model.CreateTenderReq)(*model.CreateTenderResp, error){
	resp, err := S.Storage.Client().CreateTender(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Tender ma'lumotlari databazaga saqlanmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func(S *Service) GetAllTenders(req *model.GetAllTendersReq)(*model.GetAllTendersResp, error){
	resp, err := S.Storage.Client().GetAllTenders(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Databazadan tenderlarni olib bo'lmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func(S *Service) UpdateTender(req *model.UpdateTenderReq)(*model.UpdateTenderResp, error){
	resp, err := S.Storage.Client().UpdateTender(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Tender ma'lumotlari databazada yangilanmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func(S *Service) DeleteTender(req *model.DeleteTenderReq)(*model.DeleteTenderResp, error){
	resp, err := S.Storage.Client().DeleteTender(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Tender ma'lumotlari databazadan o'chirilmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func(S *Service) GetTenderBids(req *model.GetTenderBidsReq)(*model.GetTenderBidsResp, error){
	resp, err := S.Storage.Client().GetTenderBids(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Tender bidlarini databazadan olib bo'lmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func(S *Service) BidAwarded(req *model.BidAwardedReq)(*model.BidAwardedResp, error){
	resp, err := S.Storage.Client().BidAwarded(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Bidlarga status berib bo'lmadi: %v", err))
		return nil, err
	}

	return resp, nil
}