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

// @Summary      Create bid
// @Description  Contractor can create bid to teender
// @Tags         Contractor
// @Accept       json
// @Produce      json
// @Security 		Bearer
// @Param        body body model.CreateBid true "bid infos (DeliveryTime format: dd-mm-yyyy)"
// @Param        id path string true "Tender id"
// @Success      200 {object} string "success"
// @Failure      400 {object} model.Error "error"
// @Failure      400 {object} model.Error "Server xatosi yoki CreateBid funksiyasi ishlamadi"
// @Router       /tenders/{id}/bids [post]
func (h *Handler) CreateBid(c *gin.Context) {
	req := model.CreateBid{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "error: " + err.Error()})
		return
	}
	prid := c.Param("id")
	if len(prid) == 0 {
		c.JSON(http.StatusBadRequest, model.Error{Message: "error: id is required"})
		h.Log.Error("Product ID is required")
		return
	}
	userId := c.GetString("UserID")

	resp, err := h.Service.CreateBid(&model.CreateBidInput{
		TenderID:     prid,
		ContractorID: userId,
		Price:        req.Price,
		Comments:     req.Comments,
		DeliveryTime: req.DeliveryTime,
	})

	if err != nil {
		h.Log.Error(fmt.Sprintf("CreateTender request error: %v", err))
		c.JSON(400, model.Error{Message: "CreateBid funksiyasi ishlamadi: " + err.Error()})
		return
	}

	idb, err := h.Storage.Client().GetUserByTebderId(prid)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
	}
	err = h.CreateNotification(idb, "someone is bid you tender", "someone bid you", prid)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
	}

	c.JSON(http.StatusOK, gin.H{"success": resp})
}

// @Summary      Get bids of tender
// @Description  Contractor can get all bids of a tender with optional filters
// @Tags         Contractor
// @Accept       json
// @Produce      json
// @Security 		Bearer
// @Param        id path string true "Tender ID"
// @Param        max_price query float64 false "Maximum price filter"
// @Param        max_delivery_time query string false "Maximum delivery time filter (ISO8601 format)"
// @Success      200 {object} []model.Bid "List of bids"
// @Failure      400 {object} model.Error "Bad request error"
// @Failure      400 {object} model.Error "Internal server error"
// @Router       /tenders/{id}/bids [get]
func (h *Handler) GetBidsOfTender(c *gin.Context) {
	tenderID := c.Param("id")
	if len(tenderID) == 0 {
		h.Log.Error("Tender ID is required")
		c.JSON(http.StatusBadRequest, model.Error{Message: "Tender ID is required"})
		return
	}

	maxPrice, maxPriceValue := c.Query("max_price"), 0.0
	if len(maxPrice) > 0 {
		var err error
		maxPriceValue, err = strconv.ParseFloat(maxPrice, 64)
		if err != nil {
			h.Log.Error(fmt.Sprintf("Invalid max_price value: %v", err))
			c.JSON(http.StatusBadRequest, model.Error{Message: "Invalid max_price value"})
			return
		}
	}

	maxDeliveryTime := c.Query("max_delivery_time")

	// Cache kalitini yaratish
	cacheKey := fmt.Sprintf("bids:tender_id:%s:max_price:%f:max_delivery_time:%s", tenderID, maxPriceValue, maxDeliveryTime)

	// Cache’dan tekshirish
	cachedBids, err := h.Storage.Caching().GetCache(cacheKey)
	if err == nil && cachedBids != "" {
		// Cache’da mavjud bo‘lsa, uni qaytaramiz
		c.JSON(http.StatusOK, gin.H{"bids": cachedBids})
		return
	}

	// Bazadan olish uchun parametrlar
	input := model.GetBidsInput{
		TenderID:        tenderID,
		MaxPrice:        maxPriceValue,
		MaxDeliveryTime: maxDeliveryTime,
	}

	// Bazadan ma’lumot olish
	bids, err := h.Service.GetBidsForTenderWithFilters(&input)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to fetch bids for tender %s: %v", tenderID, err))
		c.JSON(400, model.Error{Message: "Failed to fetch bids"})
		return
	}

	// Ma’lumotni JSON formatida stringga aylantirish
	bidsBytes, err := json.Marshal(bids)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Marshalling error: %v", err))
		c.JSON(400, model.Error{Message: "Failed to process bids"})
		return
	}

	// Cache’ga saqlash (10 daqiqa davomida)
	cacheErr := h.Storage.Caching().SetCache(cacheKey, string(bidsBytes), 10*time.Minute)
	if cacheErr != nil {
		h.Log.Error(fmt.Sprintf("Failed to save to cache: %v", cacheErr))
	}

	// Javobni qaytarish
	c.JSON(http.StatusOK, gin.H{"bids": bids})
}

// @Summary      Get tenders by filters
// @Description  Retrieve a list of tenders filtered by status or other parameters
// @Tags         Contractor
// @Accept       json
// @Produce      json
// @Security 		Bearer
// @Param        status query string false "Filter by status (e.g., open, closed, awarded)"
// @Success      200 {array} model.Tender "List of tenders"
// @Failure      400 {object} model.Error "Invalid request"
// @Failure      400 {object} model.Error "Server error"
// @Router       /tenders/all [get]
func (h *Handler) GetTendersByFilters(c *gin.Context) {
	// Kiruvchi ma'lumotlarni olish
	var input model.GetTendersInput
	if err := c.ShouldBindQuery(&input); err != nil {
		h.Log.Error(fmt.Sprintf("Invalid query parameters: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Invalid query parameters: " + err.Error()})
		return
	}

	// Cache uchun unikal kalit yaratish
	cacheKey := fmt.Sprintf("tenders:filters:%s",
		input.Status)

	// Cache’dan tekshirish
	cachedTenders, err := h.Storage.Caching().GetCache(cacheKey)
	if err == nil && cachedTenders != "" {
		// Agar cache’da mavjud bo‘lsa, uni qaytaramiz
		c.JSON(http.StatusOK, gin.H{"tenders": cachedTenders})
		return
	}

	// Bazadan ma'lumot olish
	tenders, err := h.Service.GetTendersByFilters(&input)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to fetch tenders: %v", err))
		c.JSON(400, model.Error{Message: "Failed to fetch tenders: " + err.Error()})
		return
	}

	// JSON formatida stringga aylantirish
	tendersBytes, err := json.Marshal(tenders)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Marshalling error: %v", err))
		c.JSON(400, model.Error{Message: "Failed to process tenders"})
		return
	}

	// Cache’ga saqlash (10 daqiqa davomida)
	cacheErr := h.Storage.Caching().SetCache(cacheKey, string(tendersBytes), 10*time.Minute)
	if cacheErr != nil {
		h.Log.Error(fmt.Sprintf("Failed to save to cache: %v", cacheErr))
	}

	// Javobni qaytarish
	c.JSON(http.StatusOK, gin.H{"tenders": tenders})
}

// @Summary      Get Bid History for a Contractor
// @Description  Retrieve all bids placed by a contractor for various tenders, including tender details like title and deadline
// @Tags         Contractor
// @Accept       json
// @Produce      json
// @Security 		Bearer
// @Param        id path string true "User ID"
// @Success      200 {array} model.BidHistory "List of bid history"
// @Failure      400 {object} model.Error "Invalid input"
// @Failure      400 {object} model.Error "Failed to retrieve bid history"
// @Router       /users/{id}/bids [get]
func (h *Handler) GetMyBidHistory(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, model.Error{Message: "user_id is required"})
		h.Log.Error("user_id is required")
		return
	}

	// Pass the pointer to the GetMyBidsInput struct
	bidHistory, err := h.Service.GetMyBidHistory(&model.GetMyBidsInput{UserID: userID})
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetMyBidHistory request error: %v", err))
		c.JSON(400, model.Error{Message: "Failed to retrieve bid history: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, bidHistory)
}
