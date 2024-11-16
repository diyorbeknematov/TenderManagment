package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"tender/model"

	"github.com/gin-gonic/gin"
)

func(h *Handler) CreateTender(c *gin.Context){
	req := model.CreateTenderReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return 
	}

	resp, err := h.Service.CreateTender(&req)
	if err != nil{
		h.Log.Error("CreateTender request error: %v", err)
		c.JSON(http.StatusInternalServerError, model.Error{Message: "CreateTender funksiyasi ishlamadi: " + err.Error()})
		return 
	}

	c.JSON(http.StatusOK, resp)
}

func(h *Handler) UpdateTender(c *gin.Context){
	req := model.UpdateTenderReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Ma'lumotlarni olishda xatolik: " + err.Error()})
		return 
	}

	resp, err := h.Service.UpdateTender(&req)
	if err != nil{
		h.Log.Error("UpdateTender request error: %v", err)
		c.JSON(http.StatusInternalServerError, model.Error{Message: "UpdateTender funksiyasi ishlamadi: " + err.Error()})
		return 
	}

	c.JSON(http.StatusOK, resp)
}

func(h *Handler) DeleteTender(c *gin.Context){
	resp, err := h.Service.DeleteTender(&model.DeleteTenderReq{Id: c.Param("id")})
	if err != nil{
		h.Log.Error("DeleteTender request error: %v", err)
		c.JSON(http.StatusInternalServerError, model.Error{Message: "DeleteTender funksiyasi ishlamadi: " + err.Error()})
		return 
	}

	c.JSON(http.StatusOK, resp)
}

func(h *Handler) GetAllTenders(c *gin.Context){
	req := model.GetAllTendersReq{}
	req.ClientId = c.Query("client_id")
	var limit, page int

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil{
		limit = 10
	}
	page, err = strconv.Atoi(c.Query("page"))
	if err != nil{
		page = 1
	}
	req.Limit = limit
	req.Page = page

	resp, err := h.Service.GetAllTenders(&req)
	if err != nil{
		h.Log.Error("GetAllTenders request error: %v", err)
		c.JSON(http.StatusInternalServerError, model.Error{Message: "GetAllTenders funksiyasi ishlamadi: " + err.Error()})
		return 
	}

	c.JSON(http.StatusOK, resp)
}

func(h *Handler) GetTenderBids(c *gin.Context){
	req := model.GetTenderBidsReq{
		ClientId: c.Query(),
	}
	req.ClientId = c.Query("client_id")
	req.TenderId = c.Param("id")
	req.
}