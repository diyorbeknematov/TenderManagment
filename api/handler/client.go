package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tender/model"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary      Tender yaratish
// @Description  Yangi tender yaratish uchun API endpoint
// @Tags         Client
// @Accept       json
// @Produce      json
// @Security 		Bearer
// @Param        body body model.CreateTenderReqSwag true "Tender yaratish uchun talab qilinadigan ma'lumotlar"
// @Success      200 {object} model.CreateTenderResp "Tender muvaffaqiyatli yaratildi"
// @Failure      400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure      500 {object} model.Error "Server xatosi yoki CreateTender funksiyasi ishlamadi"
// @Router       /tenders [post]
func (h *Handler) CreateTender(c *gin.Context) {
	req := model.CreateTenderReqSwag{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return
	}

	resp, err := h.Service.CreateTender(&model.CreateTenderReq{
		ClientId:    c.GetString("UserID"),
		Title:       req.Title,
		Description: req.Description,
		Diadline:    req.Diadline,
		Budget:      req.Budget,
	})
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
// @Security 		Bearer
// @Param        id   path      string                 true  "Tender ID"
// @Param        body body      model.UpdateTenderReqSwag  true  "Tenderni yangilash uchun talab qilinadigan ma'lumotlar"
// @Success      200  {object}  model.UpdateTenderResp "Tender ma'lumotlari muvaffaqiyatli yangilandi"
// @Failure      400  {object}  model.Error            "Ma'lumotlarni olishda xatolik"
// @Failure      500  {object}  model.Error            "Server xatosi yoki UpdateTender funksiyasi ishlamadi"
// @Router       /tenders/{id} [put]
func (h *Handler) UpdateTender(c *gin.Context) {
	req := model.UpdateTenderReqSwag{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return
	}

	resp, err := h.Service.UpdateTender(&model.UpdateTenderReq{
		Id:          c.Param("id"),
		ClientId:    c.GetString("UserID"),
		Title:       req.Title,
		Description: req.Description,
		Diadline:    req.Diadline,
		Budget:      req.Budget,
		Status:      req.Status,
	})
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
// @Security 		Bearer
// @Param        id   path      string                 true  "Tenderning ID'si"
// @Success      200  {object}  model.DeleteTenderResp "Tender muvaffaqiyatli o'chirildi"
// @Failure      500  {object}  model.Error            "Server xatosi yoki DeleteTender funksiyasi ishlamadi"
// @Router       /tenders/{id} [delete]
func (h *Handler) DeleteTender(c *gin.Context) {
	resp, err := h.Service.DeleteTender(&model.DeleteTenderReq{
		Id:       c.Param("id"),
		ClientId: c.GetString("UserID"),
	})
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
// @Security 		Bearer
// @Param        limit      query   int    false "Bir sahifadagi tenderlar soni (standart: 10)"
// @Param        page       query   int    false "Sahifa raqami (standart: 1)"
// @Success      200        {object} model.GetAllTendersResp "Tenderlar muvaffaqiyatli qaytarildi"
// @Failure      400        {object} model.Error            "Noto'g'ri parametrlar"
// @Failure      500        {object} model.Error            "Server xatosi yoki GetAllTenders funksiyasi ishlamadi"
// @Router       /tenders [get]
func (h *Handler) GetAllTenders(c *gin.Context) {
	var limit, page int

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	clientID := c.GetString("UserID")
	cacheKey := fmt.Sprintf("all_tenders:%s:%d:%d", clientID, limit, page)

	cachedResponse, err := h.Storage.Caching().GetCache(cacheKey)
	if err == nil && cachedResponse != "" {
		c.JSON(http.StatusOK, cachedResponse)
		return
	}

	resp, err := h.Service.GetAllTenders(&model.GetAllTendersReq{
		ClientId: clientID,
		Limit:    limit,
		Page:     page,
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetAllTenders request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "GetAllTenders funksiyasi ishlamadi: " + err.Error()})
		return
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Marshalling error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Javobni marshaling qilishda xatolik"})
		return
	}

	cacheErr := h.Storage.Caching().SetCache(cacheKey, string(respBytes), 5*time.Minute)
	if cacheErr != nil {
		h.Log.Error(fmt.Sprintf("Cache saqlashda xato: %v", cacheErr))
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Tenderga bildirilgan takliflarni olish
// @Description  Client o'z tenderi uchun bildirilgan barcha takliflarni olish uchun API endpoint
// @Tags         Client
// @Accept       json
// @Produce      json
// @Security 		Bearer
// @Param        id           path     string true  "Tender ID'si"
// @Param        limit        query    int    false "Bir sahifadagi takliflar soni (standart: 10)"
// @Param        page         query    int    false "Sahifa raqami (standart: 1)"
// @Param        start_price  query    string false "Takliflarning boshlang'ich narxi (filtrlash uchun ixtiyoriy, float qiymat ko'rinishida yozilishi kerak, masalan: 100.50)"
// @Param        end_price    query    string false "Takliflarning yakuniy narxi (filtrlash uchun ixtiyoriy, float qiymat ko'rinishida yozilishi kerak, masalan: 500.75)"
// @Param        start_date   query    string false "Boshlanish sanasi (filtrlash uchun ixtiyoriy, format: YYYY-MM-DD)"
// @Param        end_date     query    string false "Tugash sanasi (filtrlash uchun ixtiyoriy, format: YYYY-MM-DD)"
// @Success      200          {object} model.GetTenderBidsResp "Takliflar muvaffaqiyatli qaytarildi"
// @Failure      400          {object} model.Error             "Ma'lumotlarni olishda xatolik"
// @Failure      500          {object} model.Error             "Server xatosi yoki GetTenderBids funksiyasi ishlamadi"
// @Router       /tenders/{id}/my/bids [get]
func (h *Handler) GetTenderBids(c *gin.Context) {
	req := model.GetTenderBidsReqSwag{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return
	}
	var limit, page int

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	startPriceStr := c.Query("start_price")
	endPriceStr := c.Query("end_price")

	startPrice, err := strconv.ParseFloat(startPriceStr, 64)
	if err != nil {
		startPrice = 0 // Standart qiymat
	}
	endPrice, err := strconv.ParseFloat(endPriceStr, 64)
	if err != nil {
		endPrice = 0 // Standart qiymat
	}

	// Cache kalitini yaratish
	clientID := c.GetString("UserID")
	tenderID := c.Param("id")
	cacheKey := fmt.Sprintf("tender_bids:%s:%s:%f:%f:%s:%s:%d:%d",
		clientID, tenderID, startPrice, endPrice, req.StartDate, req.EndDate, limit, page)

	// Cache’dan tekshirish
	cachedResponse, err := h.Storage.Caching().GetCache(cacheKey)
	if err == nil && cachedResponse != "" {
		// Agar cache’da mavjud bo'lsa, uni qaytaramiz
		c.JSON(http.StatusOK, cachedResponse)
		return
	}

	// Cache’da ma'lumot yo'q bo'lsa, bazadan olish
	resp, err := h.Service.GetTenderBids(&model.GetTenderBidsReq{
		ClientId:   clientID,
		TenderId:   tenderID,
		StartPrice: startPrice,
		EndPrice:   endPrice,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Limit:      limit,
		Page:       page,
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetTenderBids request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "GetTenderBids funksiyasi ishlamadi: " + err.Error()})
		return
	}

	// Bazadan olingan javobni JSON formatida stringga o'zgartiramiz
	respBytes, err := json.Marshal(resp)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Marshalling error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Javobni marshaling qilishda xatolik"})
		return
	}

	// Cache’ga saqlash (5 daqiqa davomida)
	cacheErr := h.Storage.Caching().SetCache(cacheKey, string(respBytes), 5*time.Minute)
	if cacheErr != nil {
		h.Log.Error(fmt.Sprintf("Cache saqlashda xato: %v", cacheErr))
	}

	// Javobni foydalanuvchiga yuboramiz
	c.JSON(http.StatusOK, resp)
}

// @Summary      Tanlangan taklifni belgilash
// @Description  Tender uchun tanlangan taklifni statusini o'zgartirish
// @Tags         Client
// @Accept       json
// @Produce      json
// @Security 		Bearer
// @Param        id   path   string              true  "Tender ID'si"
// @Param        body body   model.SubmitBitReqSwag  true  "Taklif haqida ma'lumot"
// @Success      200  {object} model.SubmitBitResp "Taklif muvaffaqiyatli belgilandi"
// @Failure      400  {object} model.Error         "Ma'lumotlarni olishda xatolik"
// @Failure      500  {object} model.Error         "Server xatosi yoki BidAwarded funksiyasi ishlamadi"
// @Router       /tenders/status_change/{id}/bids [post]
func (h *Handler) SubmitBit(c *gin.Context) {
	req := model.SubmitBitReqSwag{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return
	}

	idb, err := h.Storage.Contractor().GetUserIDByBidID(req.BidId)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
	}
	err = h.CreateNotification(idb, "look at your bid", "user bit role is changed", req.BidId)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
	}

	resp, err := h.Service.SubmitBit(&model.SubmitBitReq{
		ClientId: c.GetString("UserID"),
		TenderId: c.Param("id"),
		BidId:    req.BidId,
		Status:   req.Status,
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("SubmitBit request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "SubmitBit funksiyasi ishlamadi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Tanlangan taklifni belgilash
// @Description  Tender uchun tanlangan taklifni "awarded" sifatida belgilash
// @Tags         Client
// @Accept       json
// @Produce      json
// @Security 		Bearer
// @Param        id      path     string                 true  "Tender ID"
// @Param        bid_id  path     string                 true  "Taklif ID"
// @Success      200     {object} model.AwardTenderResp  "Taklif muvaffaqiyatli belgilandi"
// @Failure      400     {object} model.Error            "Yaroqsiz ma'lumot yoki noto'g'ri so'rov"
// @Failure      404     {object} model.Error            "Tender yoki taklif topilmadi"
// @Failure      500     {object} model.Error            "Server xatosi yoki AwardTender funksiyasi ishlamadi"
// @Router       /tenders/{id}/award/{bid_id} [post]
func (h *Handler) AwardTender(c *gin.Context) {
	resp, err := h.Service.AwardTender(&model.AwardTenderReq{
		ClientId: c.GetString("UserID"),
		TenderId: c.Param("id"),
		BidId:    c.Param("bid_id"),
	})

	if err != nil {
		h.Log.Error(fmt.Sprintf("AwardTender request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "AwardTender funksiyasi ishlamadi: " + err.Error()})
		return
	}

	idb, err := h.Storage.Contractor().GetUserIDByBidID(c.Param("bid_id"))
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
	}
	err = h.CreateNotification(idb, "your bid is awarded", "user bit role is changed", c.Param("bid_id"))
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
	}

	c.JSON(http.StatusOK, resp)
}

// GetMyTenderHistory godoc
// @Summary      Foydalanuvchi tender tarixini olish
// @Description  Foydalanuvchi (client) uchun barcha tender tarixini qaytaradi. Cache-dan foydalanadi.
// @Tags         Client
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string  true  "Foydalanuvchi ID"
// @Success      200  {object}  model.GetAllTendersResp
// @Failure      400  {object}  model.Error  "Client ID kiritilmagan"
// @Failure      500  {object}  model.Error  "Ichki xatolik yoki ma'lumot olinmadi"
// @Router       /users/{id}/tenders [get]
func (h *Handler) GetMyTenderHistory(c *gin.Context) {
	clientID := c.Param("id")
	if clientID == "" {
		h.Log.Error("Client ID is required")
		c.JSON(http.StatusBadRequest, model.Error{Message: "Client ID is required"})
		return
	}

	// Cache kalitini yaratish
	cacheKey := fmt.Sprintf("tender_history_%s", clientID)

	// Cache-ni tekshirish
	cachedData, err := h.Storage.Caching().GetCache(cacheKey)
	if err == nil && cachedData != "" {
		// Cache-dan ma'lumotni qaytarish
		h.Log.Info(fmt.Sprintf("Cache hit for key: %s", cacheKey))
		var cachedResp model.GetAllTendersResp
		if err := json.Unmarshal([]byte(cachedData), &cachedResp); err == nil {
			c.JSON(http.StatusOK, cachedResp)
			return
		} else {
			h.Log.Error(fmt.Sprintf("Failed to unmarshal cached data: %v", err))
		}
	}

	// Cache-da ma'lumot yo'q, bazadan olish
	resp, err := h.Storage.Client().GetAllTenders(&model.GetAllTendersReq{
		ClientId: clientID,
		Limit:    1000,
		Page:     1,
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetMyTenderHistory request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "GetMyTenderHistory funksiyasi ishlamadi: " + err.Error()})
		return
	}

	// Cache-ga yozish
	respData, err := json.Marshal(resp)
	if err == nil {
		if err := h.Storage.Caching().SetCache(cacheKey, string(respData), time.Hour*1); err != nil {
			h.Log.Error(fmt.Sprintf("Failed to cache data for key %s: %v", cacheKey, err))
		} else {
			h.Log.Info(fmt.Sprintf("Cached response for key: %s", cacheKey))
		}
	} else {
		h.Log.Error(fmt.Sprintf("Failed to marshal response for caching: %v", err))
	}

	// Javobni qaytarish
	c.JSON(http.StatusOK, resp)
}
