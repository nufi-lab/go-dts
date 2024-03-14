package ordercontroller

import (
	"challenge-2/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var orders []models.Order

	models.DB.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func Show(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := models.DB.First(&order, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func Create(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Set ordered_at to the current time
	order.OrderedAt = time.Now()
	models.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func Update(c *gin.Context) {
	var order models.Order

	id := c.Param("id")
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&order).Where("order_id = ?", id).Updates(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var order models.Order

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	if models.DB.Delete(&order, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})

}
