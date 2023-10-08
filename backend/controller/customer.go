package controller

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/Kami0rn/SoyJuu/entities"
)


//POST /customer
func CreateCustomer(c *gin.Context) {
	var customer entities.Customer
	
	//bind
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	// Create a new customer entity with the given customer information.
	u := entities.Customer{
		FirstName: customer.FirstName,
		LastName: customer.LastName,
		Username: customer.Username,
		Password: customer.Password,
		Address: customer.Address,
		Email: customer.Email,
		Phone: customer.Phone,
		Gender: customer.Gender,
	}

	if err := entities.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})

}