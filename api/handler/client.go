package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"tender/model"

	"github.com/gin-gonic/gin"
)

// @Summary      Tender yaratish
// @Description  Yangi tender yaratish uchun API endpoint
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        body body model.CreateTenderReq true "Tender yaratish uchun talab qilinadigan ma'lumotlar"
// @Success      200 {object} model.CreateTenderResp "Tender muvaffaqiyatli yaratildi"
// @Failure      400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure      500 {object} model.Error "Server xatosi yoki CreateTender funksiyasi ishlamadi"
// @Router       /tender/create [post]
func (h *Handler) CreateTender(c *gin.Context) {
	req := model.CreateTenderReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return
	}

	resp, err := h.Service.CreateTender(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("CreateTender request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "CreateTender funksiyasi ishlamadi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Tender ma'lumotlarini yangilash
// @Description  Mavjud tenderning ma'lumotlarini yangilash uchun API endpoint
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        body body model.UpdateTenderReq true "Tenderni yangilash uchun talab qilinadigan ma'lumotlar"
// @Success      200 {object} model.UpdateTenderResp "Tender ma'lumotlari muvaffaqiyatli yangilandi"
// @Failure      400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure      500 {object} model.Error "Server xatosi yoki UpdateTender funksiyasi ishlamadi"
// @Router       /tender/update [put]
func (h *Handler) UpdateTender(c *gin.Context) {
	req := model.UpdateTenderReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return
	}

	resp, err := h.Service.UpdateTender(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("UpdateTender request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "UpdateTender funksiyasi ishlamadi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Tenderni o'chirish
// @Description  Mavjud tenderni o'chirish uchun API endpoint
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        id path string true "Tenderning ID'si"
// @Success      200 {object} model.DeleteTenderResp "Tender muvaffaqiyatli o'chirildi"
// @Failure      500 {object} model.Error "Server xatosi yoki DeleteTender funksiyasi ishlamadi"
// @Router       /tender/delete/{id} [delete]
func (h *Handler) DeleteTender(c *gin.Context) {
	resp, err := h.Service.DeleteTender(&model.DeleteTenderReq{Id: c.Param("id")})
	if err != nil {
		h.Log.Error(fmt.Sprintf("DeleteTender request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "DeleteTender funksiyasi ishlamadi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Tenderlar ro'yxatini olish
// @Description  Tenderlar ro'yxatini olish uchun API endpoint
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        client_id query string false "Mijoz ID'si (filtrlash uchun ixtiyoriy)"
// @Param        limit query int false "Bir sahifadagi tenderlar soni (standart: 10)"
// @Param        page query int false "Sahifa raqami (standart: 1)"
// @Success      200 {object} model.GetAllTendersResp "Tenderlar muvaffaqiyatli qaytarildi"
// @Failure      500 {object} model.Error "Server xatosi yoki GetAllTenders funksiyasi ishlamadi"
// @Router       /tender/get_all [get]
func (h *Handler) GetAllTenders(c *gin.Context) {
	req := model.GetAllTendersReq{}
	req.ClientId = c.Query("client_id")
	var limit, page int

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	req.Limit = limit
	req.Page = page

	resp, err := h.Service.GetAllTenders(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetAllTenders request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "GetAllTenders funksiyasi ishlamadi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Tenderga bildirilgan takliflarni olish
// @Description  Ma'lum bir tender uchun bildirilgan barcha takliflarni olish uchun API endpoint
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        id path string true "Tender ID'si"
// @Param        limit query int false "Bir sahifadagi takliflar soni (standart: 10)"
// @Param        page query int false "Sahifa raqami (standart: 1)"
// @Success      200 {object} model.GetTenderBidsResp "Takliflar muvaffaqiyatli qaytarildi"
// @Failure      400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure      500 {object} model.Error "Server xatosi yoki GetTenderBids funksiyasi ishlamadi"
// @Router       /tender/{id}/bids [get]
func (h *Handler) GetTenderBids(c *gin.Context) {
	req := model.GetTenderBidsReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return
	}
	req.TenderId = c.Param("id")
	var limit, page int

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	req.Limit = limit
	req.Page = page

	resp, err := h.Service.GetTenderBids(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetTenderBids request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "GetTenderBids funksiyasi ishlamadi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Tanlangan taklifni belgilash
// @Description  Tender uchun tanlangan taklifni "awarded" sifatida belgilash
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param        body body model.BidAwardedReq true "Taklif haqida ma'lumot"
// @Success      200 {object} model.BidAwardedResp "Taklif muvaffaqiyatli belgilandi"
// @Failure      400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure      500 {object} model.Error "Server xatosi yoki BidAwarded funksiyasi ishlamadi"
// @Router       /tender/bid_awarded [post]
func (h *Handler) BidAwarded(c *gin.Context) {
	req := model.BidAwardedReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return
	}

	resp, err := h.Service.BidAwarded(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetTenderBids request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "GetTenderBids funksiyasi ishlamadi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
