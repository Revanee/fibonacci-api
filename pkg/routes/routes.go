package routes

import (
	"net/http"
	"strconv"

	"github.com/Revanee/fibonacci-api/pkg/numbers"
	"github.com/gin-gonic/gin"
)

func CheckFibonacci(c *gin.Context) {
	numberParam := c.Param("number")
	number, err := strconv.Atoi(numberParam)
	if err != nil {
		c.String(http.StatusBadRequest, "%v is not a number", numberParam)
		return
	}
	if numbers.IsFibonacci(number) {
		c.JSON(http.StatusOK, gin.H{
			"previous":     numbers.PreviousFibonacci(number),
			"next":         numbers.NextFibonacci(number),
			"is_fibonacci": "True",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"is_fibonacci": "False",
		})
	}
}

func CheckPrime(c *gin.Context) {
	numberParam := c.Param("number")
	number, err := strconv.Atoi(numberParam)
	if err != nil {
		c.String(http.StatusBadRequest, "%v is not a number", numberParam)
		return
	}
	if numbers.IsPrime(number) {
		c.JSON(http.StatusOK, gin.H{
			"message": "True",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "False",
		})
	}
}
