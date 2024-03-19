package ordercontroller

import (
	"assignment-3/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var requestData struct {
		OrderedAt    string `json:"orderedAt"`
		CustomerName string `json:"customerName"`
		Items        []struct {
			ItemCode    string `json:"itemCode"`
			Description string `json:"description"`
			Quantity    int    `json:"quantity"`
		} `json:"items"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Parse time
	// orderedAt, err := time.Parse("2006-01-02 15:04:05.9999", requestData.OrderedAt)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": "Invalid time format"})
	// 	return
	// }

	// Create order
	order := models.Order{
		CustomerName: requestData.CustomerName,
		// Set ordered_at to the current time
		OrderedAt: time.Now(),
		// OrderedAt:    orderedAt,
	}

	models.DB.Create(&order) // Create order first

	// Create items
	for _, itemData := range requestData.Items {
		item := models.Item{
			ItemCode:    itemData.ItemCode,
			Description: itemData.Description,
			Quantity:    itemData.Quantity,
			OrderID:     order.OrderID, // Set foreign key
		}
		models.DB.Create(&item) // Create item associated with the order
	}

	c.JSON(201, gin.H{"message": "Order created successfully"})
}

func Index(c *gin.Context) {
	var orders []models.Order
	models.DB.Preload("Items").Find(&orders)

	// Return orders as JSON response
	c.JSON(200, orders)
}

func Update(c *gin.Context) {
	// Get order ID from URL parameter
	orderID := c.Param("id")

	// Parse request body to extract updated order data
	var updatedOrderData struct {
		OrderedAt    string `json:"orderedAt"`
		CustomerName string `json:"customerName"`
		Items        []struct {
			ItemID      uint   `json:"itemID"`
			ItemCode    string `json:"itemCode"`
			Description string `json:"description"`
			Quantity    int    `json:"quantity"`
		} `json:"items"`
	}
	if err := c.BindJSON(&updatedOrderData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve existing order from the database
	var existingOrder models.Order
	if err := models.DB.Preload("Items").First(&existingOrder, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Update existing order with new data
	existingOrder.CustomerName = updatedOrderData.CustomerName
	orderedAt, err := time.Parse("2006-01-02 15:04:05.9999", updatedOrderData.OrderedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}
	existingOrder.OrderedAt = orderedAt

	// Update items associated with the order
	for _, updatedItem := range updatedOrderData.Items {
		// Find the corresponding item based on itemID and orderID
		var existingItem models.Item
		if err := models.DB.Where("order_id = ? AND item_id = ?", orderID, updatedItem.ItemID).First(&existingItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		// Update existing item with new data
		existingItem.ItemCode = updatedItem.ItemCode
		existingItem.Description = updatedItem.Description
		existingItem.Quantity = updatedItem.Quantity

		// Save updated item back to the database
		if err := models.DB.Save(&existingItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
			return
		}
	}

	// Save updated order back to the database
	if err := models.DB.Save(&existingOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, existingOrder)
}

func Delete(c *gin.Context) {
	// Get order ID from URL parameter
	orderID := c.Param("id")

	// Retrieve existing order from the database
	var existingOrder models.Order
	if err := models.DB.First(&existingOrder, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Delete related items first
	if err := models.DB.Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete related items"})
		return
	}

	// Delete the order from the database
	if err := models.DB.Delete(&existingOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Order and related items deleted successfully"})
}

// func Show(c *gin.Context) {
// 	var order models.Order
// 	id := c.Param("id")

// 	if err := models.DB.First(&order, id).Error; err != nil {
// 		switch err {
// 		case gorm.ErrRecordNotFound:
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
// 			return
// 		default:
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"order": order})
// }

// func Create(c *gin.Context) {
// 	var order models.Order

// 	if err := c.ShouldBindJSON(&order); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	// Set ordered_at to the current time
// 	order.OrderedAt = time.Now()
// 	models.DB.Create(&order)
// 	c.JSON(http.StatusOK, gin.H{"order": order})
// }

// func Update(c *gin.Context) {
// 	var order models.Order

// 	id := c.Param("id")
// 	if err := c.ShouldBindJSON(&order); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	if models.DB.Model(&order).Where("order_id = ?", id).Updates(&order).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate order"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
// }

// func Delete(c *gin.Context) {
// 	var order models.Order

// 	var input struct {
// 		Id json.Number
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	id, _ := input.Id.Int64()

// 	if models.DB.Delete(&order, id).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus order"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})

// }
