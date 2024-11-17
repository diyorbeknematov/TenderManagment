package service

import (
	"fmt"
	"tender/model"
	"time"
)

func (S *Service) CreateTender(req *model.CreateTenderReq) (*model.CreateTenderResp, error) {
	if req.Budget < 0 {
		S.Log.Error("Budget 0 dan kichik bo'lishi mumkin emas")
		return nil, fmt.Errorf("budget 0 dan kichik bo'lishi mumkin emas")
	}
	if len(req.Diadline) != 0 && time.Now().String() > req.Diadline {
		S.Log.Error("Diadline hozirgi vaqtdan ortda bo'lishi mumkin emas")
		return nil, fmt.Errorf("diadline hozirgi vaqtdan ortda bo'lishi mumkin emas")
	}

	resp, err := S.Storage.Client().CreateTender(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Tender ma'lumotlari databazaga saqlanmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func (S *Service) GetAllTenders(req *model.GetAllTendersReq) (*model.GetAllTendersResp, error) {
	resp, err := S.Storage.Client().GetAllTenders(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Databazadan tenderlarni olib bo'lmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func (S *Service) UpdateTender(req *model.UpdateTenderReq) (*model.UpdateTenderResp, error) {
	if req.Budget < 0 {
		S.Log.Error("Budget 0 dan kichik bo'lishi mumkin emas")
		return nil, fmt.Errorf("budget 0 dan kichik bo'lishi mumkin emas")
	}
	if time.Now().String() > req.Diadline {
		S.Log.Error("Diadline hozirgi vaqtdan ortda bo'lishi mumkin emas")
		return nil, fmt.Errorf("diadline hozirgi vaqtdan ortda bo'lishi mumkin emas")
	}

	resp, err := S.Storage.Client().UpdateTender(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Tender ma'lumotlari databazada yangilanmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func (S *Service) DeleteTender(req *model.DeleteTenderReq) (*model.DeleteTenderResp, error) {
	resp, err := S.Storage.Client().DeleteTender(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Tender ma'lumotlari databazadan o'chirilmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func (S *Service) GetTenderBids(req *model.GetTenderBidsReq) (*model.GetTenderBidsResp, error) {
	if req.StartPrice < 0 || req.EndPrice < 0 {
		S.Log.Error("Deadline hozirgi vaqtdan ortda bo'lishi mumkin emas")
		return nil, fmt.Errorf("budget 0 dan kichik bo'lishi mumkin emas")
	}
	clientId, err := S.Storage.Client().GetUserByTebderId(req.TenderId)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Bunday tender mavjud emas: %v", err))
		return nil, err
	}
	if clientId != req.ClientId {
		S.Log.Error(fmt.Sprintf("Clientning bunday tenderi mavjud emas: %v", err))
		return nil, fmt.Errorf("clientning bunday tenderi mavjud emas: %v", err)
	}

	resp, err := S.Storage.Client().GetTenderBids(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Tender bidlarini databazadan olib bo'lmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func (S *Service) SubmitBit(req *model.SubmitBitReq) (*model.SubmitBitResp, error) {
	clientId, err := S.Storage.Client().GetUserByTebderId(req.TenderId)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Bunday tender mavjud emas: %v", err))
		return nil, err
	}
	if clientId != req.ClientId {
		S.Log.Error(fmt.Sprintf("Clientning bunday tenderi mavjud emas: %v", err))
		return nil, fmt.Errorf("clientning bunday tenderi mavjud emas: %v", err)
	}

	resp, err := S.Storage.Client().SubmitBit(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Bidlarga status berib bo'lmadi: %v", err))
		return nil, err
	}

	return resp, nil
}

func (S *Service) AwardTender(req *model.AwardTenderReq) (*model.AwardTenderResp, error) {
	resp, err := S.Storage.Client().AwardTender(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Tender yakunlanmadi: %v", err))
		return nil, err
	}

	return resp, nil
}
