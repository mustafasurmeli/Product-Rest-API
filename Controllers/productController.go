package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trainRestApi/Database"
	"trainRestApi/Models"
)

func GetProducts(c *gin.Context) {
	var products []Models.Product
	rows, err := Database.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}
	c.JSON(http.StatusOK, products)

}

func CreateProduct(c *gin.Context) {
	var product Models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id"
	err := Database.DB.QueryRow(query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	var product Models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := "UPDATE products SET name=$1, price=$2 WHERE id=$3"
	_, err := Database.DB.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)

}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	query := "DELETE FROM products WHERE id=$1"
	_, err := Database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
