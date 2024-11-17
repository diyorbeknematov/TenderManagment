package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"tender/model"

	"github.com/gin-gonic/gin"
)

// @Summary      Create bid
// @Description  Contractor can create bid to teender
// @Tags         Contractor
// @Accept       json
// @Produce      json
// @Param        body body model.CreateBid true "bid infos (DeliveryTime format: dd-mm-yyyy)"
// @Param        id path string true "Tender id"
// @Success      200 {object} string "success"
// @Failure      400 {object} model.Error "error"
// @Failure      500 {object} model.Error "Server xatosi yoki CreateBid funksiyasi ishlamadi"
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
		c.JSON(http.StatusInternalServerError, model.Error{Message: "CreateBid funksiyasi ishlamadi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp})
}

// @Summary      Get bids of tender
// @Description  Contractor can get all bids of a tender with optional filters
// @Tags         Contractor
// @Accept       json
// @Produce      json
// @Param        id path string true "Tender ID"
// @Param        max_price query float64 false "Maximum price filter"
// @Param        max_delivery_time query string false "Maximum delivery time filter (ISO8601 format)"
// @Success      200 {object} []model.Bid "List of bids"
// @Failure      400 {object} model.Error "Bad request error"
// @Failure      500 {object} model.Error "Internal server error"
// @Router       /tenders/{id}/bids [get]
func (h *Handler) GetBidsOfTender(c *gin.Context) {
	tenderID := c.Param("id")
	if len(tenderID) == 0 {
		h.Log.Error("Tender ID is required")
		c.JSON(http.StatusBadRequest, model.Error{Message: "Tender ID is required"})
		return
	}

	maxPrice, maxPriceErr := c.Query("max_price"), 0.0
	if len(maxPrice) > 0 {
		var err error
		maxPriceErr, err = strconv.ParseFloat(maxPrice, 64)
		if err != nil {
			h.Log.Error(fmt.Sprintf("Invalid max_price value: %v", err))
			c.JSON(http.StatusBadRequest, model.Error{Message: "Invalid max_price value"})
			return
		}
	}

	maxDeliveryTime := c.Query("max_delivery_time")

	input := model.GetBidsInput{
		TenderID:        tenderID,
		MaxPrice:        maxPriceErr,
		MaxDeliveryTime: maxDeliveryTime,
	}

	bids, err := h.Service.GetBidsForTenderWithFilters(&input)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to fetch bids for tender %s: %v", tenderID, err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Failed to fetch bids"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bids": bids})
}

// @Summary      Get tenders by filters
// @Description  Retrieve a list of tenders filtered by status or other parameters
// @Tags         Tenders
// @Accept       json
// @Produce      json
// @Param        status query string false "Filter by status (e.g., open, closed)"
// @Success      200 {array} model.Tender "List of tenders"
// @Failure      400 {object} model.Error "Invalid request"
// @Failure      500 {object} model.Error "Server error"
// @Router       /tenders/all [get]
func (h *Handler) GetTendersByFilters(c *gin.Context) {
	var input model.GetTendersInput
	if err := c.ShouldBindQuery(&input); err != nil {
		h.Log.Error(fmt.Sprintf("Invalid query parameters: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Invalid query parameters: " + err.Error()})
		return
	}

	tenders, err := h.Service.GetTendersByFilters(&input)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to fetch tenders: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Failed to fetch tenders: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenders)
}
