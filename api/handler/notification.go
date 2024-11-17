package handler

// func (h *Handler) SendNotification(c *gin.Context) {
// 	// userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
// 		return
// 	}

// 	message := c.PostForm("message")
// 	if message == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "message is required"})
// 		return
// 	}

// 	// if err := h.Service.SendNotification(userID, message); err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send notification"})
// 	// 	return
// 	// }

// 	c.JSON(http.StatusOK, gin.H{"status": "notification sent"})
// }
